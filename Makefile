include bin/build/make/client.mak
include bin/build/make/git.mak

# Create a pipeline.
create:
	(cd test &&../idpctl pipeline -i file:.config/idpd.yml --create pipeline)
