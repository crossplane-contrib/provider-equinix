module github.com/crossplane-contrib/provider-equinix

go 1.16

require (
	github.com/crossplane/crossplane-runtime v0.13.0
	github.com/crossplane/crossplane-tools v0.0.0-20210320162312-1baca298c527
	github.com/equinix/ne-go v1.1.0
	github.com/equinix/oauth2-go v1.0.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.9.1
	golang.org/x/tools v0.0.0-20200916195026-c9a70fc28ce3 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
	k8s.io/apimachinery v0.20.1
	k8s.io/client-go v0.20.1
	sigs.k8s.io/controller-runtime v0.8.0
	sigs.k8s.io/controller-tools v0.3.0
)
