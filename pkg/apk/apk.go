package apk

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"

	"github.com/chainguard-dev/go-apk/pkg/apk"
	"github.com/go-ini/ini"
	"github.com/cvedb/cvectl/pkg/versions"
)

type Context struct {
	client   *http.Client
	indexURL string
}

func New(client *http.Client, indexURL string) Context {
	return Context{
		client:   client,
		indexURL: indexURL,
	}
}

func (c Context) GetApkPackages() (map[string]*apk.Package, error) {
	req, err := http.NewRequest("GET", c.indexURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting URI %s: %w", c.indexURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non ok http response for URI %s code: %v", c.indexURL, resp.StatusCode)
	}

	return ParseApkIndex(resp.Body)
}

func ParseUnpackedApkIndex(indexData io.ReadCloser) (map[string]*apk.Package, error) {
	cvedbPackages := make(map[string]*apk.Package)

	packages, err := apk.ParsePackageIndex(indexData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response %v: %w", indexData, err)
	}

	return getLatestPackagesMap(packages, cvedbPackages)
}

func getLatestPackagesMap(apkIndexPackages []*apk.Package, cvedbPackages map[string]*apk.Package) (map[string]*apk.Package, error) {
	for _, p := range apkIndexPackages {
		if cvedbPackages[p.Name] != nil {
			vers := []string{cvedbPackages[p.Name].Version, p.Version}
			sort.Sort(versions.ByLatestStrings(vers))
			// replace in our map if we find a newer version in the APKINDEX
			if p.Version == vers[0] {
				cvedbPackages[p.Name] = p
			}
		} else {
			cvedbPackages[p.Name] = p
		}
	}
	log.Printf("found %d latest apk index package versions", len(cvedbPackages))
	return cvedbPackages, nil
}

func ParseApkIndex(indexData io.ReadCloser) (map[string]*apk.Package, error) {
	cvedbPackages := make(map[string]*apk.Package)

	apkIndex, err := apk.IndexFromArchive(indexData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response %v: %w", indexData, err)
	}

	return getLatestPackagesMap(apkIndex.Packages, cvedbPackages)
}

func PKGINFOFromAPK(r io.Reader) (*apk.Package, error) {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}

	tr := tar.NewReader(gr)
	var pkginfoHdr *tar.Header
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read tar header: %w", err)
		}

		if hdr.Name == ".PKGINFO" {
			pkginfoHdr = hdr
			break
		}
	}

	if pkginfoHdr == nil {
		return nil, fmt.Errorf("failed to find .PKGINFO in apk")
	}

	pkginfo := new(apk.Package)
	loaded, err := ini.Load(tr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse INI data: %w", err)
	}
	err = loaded.MapTo(pkginfo)
	if err != nil {
		return nil, fmt.Errorf("failed to map INI data to struct: %w", err)
	}

	return pkginfo, nil
}
