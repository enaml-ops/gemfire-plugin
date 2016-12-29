package server 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Gemfire struct {

	/*Locator - Descr: Port the Locator is listening on Default: <nil>
*/
	Locator *Locator `yaml:"locator,omitempty"`

	/*Authn - Descr: base64 encoding of authentication security jar Default: 
*/
	Authn *Authn `yaml:"authn,omitempty"`

	/*ClusterTopology - Descr: Number of locators per cluster Default: 0
*/
	ClusterTopology *ClusterTopology `yaml:"cluster_topology,omitempty"`

	/*Server - Descr: Port for the REST endpoint for administration Default: 8080
*/
	Server *Server `yaml:"server,omitempty"`

}