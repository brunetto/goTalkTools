goTalkTools
===========

Simple and raw tool to prepare and manage a beamer presentation from a template with go(lang). 
It produces  as output three different presentation:    

1. - normal presentation with appendix
2. - presentation without appendix for sharing
3. - handout notes
It also copy the sharable presentation to the (assumed) publicfolder for upload.


Note: .tex file is assumed to be in the provided template format

### Usage    

````bash
  gotalk
  gotalk [command]

Available Commands: 
  version         :: Print the version number of slt
  prepare         :: Prepare a new folder from template to start a presentation
  compile         :: Compile the presentation.
  publish         :: Publish the sharable version of the presentation.
  help [command]  :: Help about any command


Use "gotalk help [command]" for more information about that command.
````
