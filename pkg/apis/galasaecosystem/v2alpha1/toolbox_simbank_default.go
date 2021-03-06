/*
 * Copyright contributors to the Galasa Project
 */
package v2alpha1

func DefaultSimbank(c *GalasaEcosystemSpec) ComponentSpec {
	return ComponentSpec{
		Image:           SIMBANKIMAGE,
		Replicas:        &SINGLEREPLICA,
		ImagePullPolicy: c.ImagePullPolicy,
		NodeSelector:    c.NodeSelector,
	}
}

func SetSimbankDefaults(simbank ComponentSpec, c *GalasaEcosystemSpec) ComponentSpec {
	if simbank.Image == "" {
		simbank.Image = SIMBANKIMAGE
	}
	if simbank.ImagePullPolicy == "" {
		simbank.ImagePullPolicy = c.ImagePullPolicy
	}
	if simbank.Replicas == nil {
		simbank.Replicas = &SINGLEREPLICA
	}
	if simbank.NodeSelector == nil {
		simbank.NodeSelector = c.NodeSelector
	}

	return simbank
}

var DEFAULTSIMBANKCATALOG = `{
	"classes": {
	  "dev.galasa.simbank.tests/dev.galasa.simbank.tests.SimBankIVT": {
		"name": "dev.galasa.simbank.tests.SimBankIVT",
		"bundle": "dev.galasa.simbank.tests",
		"shortName": "SimBankIVT",
		"package": "dev.galasa.simbank.tests"
	  },
	  "dev.galasa.simbank.tests/dev.galasa.simbank.tests.BatchAccountsOpenTest": {
		"name": "dev.galasa.simbank.tests.BatchAccountsOpenTest",
		"bundle": "dev.galasa.simbank.tests",
		"shortName": "BatchAccountsOpenTest",
		"package": "dev.galasa.simbank.tests"
	  },
	  "dev.galasa.simbank.tests/dev.galasa.simbank.tests.BasicAccountCreditTest": {
		"name": "dev.galasa.simbank.tests.BasicAccountCreditTest",
		"bundle": "dev.galasa.simbank.tests",
		"shortName": "BasicAccountCreditTest",
		"package": "dev.galasa.simbank.tests"
	  },
	  "dev.galasa.simbank.tests/dev.galasa.simbank.tests.ProvisionedAccountCreditTests": {
		"name": "dev.galasa.simbank.tests.ProvisionedAccountCreditTests",
		"bundle": "dev.galasa.simbank.tests",
		"shortName": "ProvisionedAccountCreditTests",
		"package": "dev.galasa.simbank.tests"
	  }
	},
	"packages": {
	  "dev.galasa.simbank.tests": [
		"dev.galasa.simbank.tests/dev.galasa.simbank.tests.SimBankIVT",
		"dev.galasa.simbank.tests/dev.galasa.simbank.tests.BatchAccountsOpenTest",
		"dev.galasa.simbank.tests/dev.galasa.simbank.tests.BasicAccountCreditTest",
		"dev.galasa.simbank.tests/dev.galasa.simbank.tests.ProvisionedAccountCreditTests"
	  ]
	},
	"bundles": {
	  "dev.galasa.simbank.tests": {
		"packages": {
		  "dev.galasa.simbank.tests": [
			"dev.galasa.simbank.tests/dev.galasa.simbank.tests.SimBankIVT",
			"dev.galasa.simbank.tests/dev.galasa.simbank.tests.BatchAccountsOpenTest",
			"dev.galasa.simbank.tests/dev.galasa.simbank.tests.BasicAccountCreditTest",
			"dev.galasa.simbank.tests/dev.galasa.simbank.tests.ProvisionedAccountCreditTests"
		  ]
		}
	  },
	  "dev.galasa.simbank.manager": {
		"packages": {}
	  }
	},
	"sharedEnvironments": {},
	"name": "Galasa SimBank OBR",
	"build": "dev.galasa:dev.galasa.simbank.obr:0.15.0 - 2020-06-10T05:32:31.267Z",
	"version": "0.15.0",
	"built": "2020-06-10T05:32:31.267Z"
  }`
