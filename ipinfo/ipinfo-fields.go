// Code generated by gen-fields; DO NOT EDIT.

package ipinfo

import (
	"bytes"
	"net"
	"strings"
)

// GetHostname returns a specific field "hostname" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetHostname(ip net.IP) (string, error) {
	return c.GetHostname(ip)
}

// GetHostname returns a specific field "hostname" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetHostname(ip net.IP) (string, error) {
	s := "hostname"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	vs := strings.TrimSpace(v.String())
	if vs == "undefined" {
		vs = ""
	}
	return vs, nil
}

// GetOrganization returns a specific field "org" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetOrganization(ip net.IP) (string, error) {
	return c.GetOrganization(ip)
}

// GetOrganization returns a specific field "org" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetOrganization(ip net.IP) (string, error) {
	s := "org"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	vs := strings.TrimSpace(v.String())
	if vs == "undefined" {
		vs = ""
	}
	return vs, nil
}

// GetCity returns a specific field "city" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetCity(ip net.IP) (string, error) {
	return c.GetCity(ip)
}

// GetCity returns a specific field "city" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetCity(ip net.IP) (string, error) {
	s := "city"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	vs := strings.TrimSpace(v.String())
	if vs == "undefined" {
		vs = ""
	}
	return vs, nil
}

// GetRegion returns a specific field "region" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetRegion(ip net.IP) (string, error) {
	return c.GetRegion(ip)
}

// GetRegion returns a specific field "region" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetRegion(ip net.IP) (string, error) {
	s := "region"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	vs := strings.TrimSpace(v.String())
	if vs == "undefined" {
		vs = ""
	}
	return vs, nil
}

// GetCountry returns a specific field "country" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetCountry(ip net.IP) (string, error) {
	return c.GetCountry(ip)
}

// GetCountry returns a specific field "country" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetCountry(ip net.IP) (string, error) {
	s := "country"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	vs := strings.TrimSpace(v.String())
	if vs == "undefined" {
		vs = ""
	}
	return vs, nil
}

// GetLocation returns a specific field "loc" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetLocation(ip net.IP) (string, error) {
	return c.GetLocation(ip)
}

// GetLocation returns a specific field "loc" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetLocation(ip net.IP) (string, error) {
	s := "loc"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	vs := strings.TrimSpace(v.String())
	if vs == "undefined" {
		vs = ""
	}
	return vs, nil
}

// GetPhone returns a specific field "phone" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetPhone(ip net.IP) (string, error) {
	return c.GetPhone(ip)
}

// GetPhone returns a specific field "phone" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetPhone(ip net.IP) (string, error) {
	s := "phone"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	vs := strings.TrimSpace(v.String())
	if vs == "undefined" {
		vs = ""
	}
	return vs, nil
}

// GetPostal returns a specific field "postal" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetPostal(ip net.IP) (string, error) {
	return c.GetPostal(ip)
}

// GetPostal returns a specific field "postal" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetPostal(ip net.IP) (string, error) {
	s := "postal"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	req, err := c.NewRequest(s)
	if err != nil {
		return "", err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return "", err
	}
	vs := strings.TrimSpace(v.String())
	if vs == "undefined" {
		vs = ""
	}
	return vs, nil
}
