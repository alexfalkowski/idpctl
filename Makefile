include bin/build/make/client.mak
include bin/build/make/git.mak

# Create a pipeline.
create:
	(cd test &&../idpctl pipeline -i file:.config/idpd.yml --create $(pipeline))

# Get a pipeline.
get:
	(cd test &&../idpctl pipeline -i file:.config/idpd.yml --get $(id))

# Update a pipeline.
update:
	(cd test &&../idpctl pipeline -i file:.config/idpd.yml --update $(id):$(pipeline))

# Delete a pipeline.
delete:
	(cd test &&../idpctl pipeline -i file:.config/idpd.yml --delete $(id))

# Trigger a pipeline.
trigger:
	(cd test &&../idpctl pipeline -i file:.config/idpd.yml --trigger $(id))
