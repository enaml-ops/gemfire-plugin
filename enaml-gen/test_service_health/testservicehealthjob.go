package test_service_health 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type TestServiceHealthJob struct {

	/*Cc - Descr: Cloud Controller Application Domain Default: <nil>
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*Gemfire - Descr: Name to give test GemFire CF service Default: gemfire-smoke-test-service-70552f90-7785-4347-9516
*/
	Gemfire *Gemfire `yaml:"gemfire,omitempty"`

	/*Broker - Descr: Gemfire service broker API address Default: <nil>
*/
	Broker *Broker `yaml:"broker,omitempty"`

}