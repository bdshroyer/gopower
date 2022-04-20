# gopower

Basic RAPL-based power monitor. WIP.

### Important Dependencies:
* [Gomega](https://github.com/onsi/gomega): Matchers for testing.
* [Ginkgo](https://github.com/onsi/ginkgo): BDD testing framework for Go.
* [Intel PowerGadget](https://www.intel.com/content/www/us/en/developer/articles/tool/power-gadget.html) *(Mac OS only)*: Provides access to Intel's power performance counter system.

### OS X Requirements

OS X doesn't have a native Intel RAPL driver, so you will need to install the [Intel PowerGadget](https://www.intel.com/content/www/us/en/developer/articles/tool/power-gadget.html) to make this work on a Mac.
