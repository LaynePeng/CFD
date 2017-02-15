# What is it?
CFD is short for Composable Function Discoverer, which is a simple tool to discover what hardware functions supported in host. Currently, it is the initial stage, only support a handful feature to tell what:

* GPU
* Intel QAT Card
* NVMe 

In each hardware function, it supported to:

* List all hardware function supported in the host;
* Test if this hardware function support;
* Show the description of the hardware function. 

We have to know this project is just begin and provide a extensible design for future implementations. 

# Limitation
Please note CFD is only tested in CentOS7 and Ubuntu 14.04, it may have problem when runned in other OS. But it provides an easy extensible design for implement other OS

# Contribution Guide
Any contribution is welcome and nessary to this project. But before contributing, please note the following terms:

* Sensor: the logic how to test the feature is supported in the host. Add a new OS supprot, please follow the file name convention: sensor_NameOfOS. Adding a new hardware function support is added to Sensor interface currently. But I will add the pluggable support in days;
* Spec: very important in the project, which will define the standard of the interfaces. It includes:
    * the Type define how to descript a hardware functions. Please note that the new hardware functions added can use import mechanism in Golang;
    * the common functions for spec.
