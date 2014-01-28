package got

import (
	"github.com/spf13/cobra"
	"fmt"
)

var GoTalkCmd = &cobra.Command{
	Use:   "gotalk",
	Short: "Prepare presentations in (for now) beamer",
	Long: `This program take a .tex beamer presentation and
produces as output three different presentation:
1 - normal presentation with appendix
2 - presentation without appendix for sharing
3 - handout notes
.tex file need to be in the brunetto's format`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Choose a sub-command or type gotalk help for help.")
	},
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of slt",
	Long:  `All software has versions. This is sltools'`,
	Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("GoTalk Tools v0.1")
	},
}

var (
	presentationName string
)

var PrepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare a new folder from template to start a presentation",
	Long:  `Download the beamer template and rename the folder according to the 
provided name for the presentation provided with the -p flag.
It would be good that the name is in the form:

<date>-<place>-<event>

for example:
	
2000-01-01-Venice-New_Years_Eve
`,
	Run: func(cmd *cobra.Command, args []string) {
		Prepare(presentationName)
	},
}


var (
	texName string
)

var CompileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile the presentation.",
	Long:  `Compile the presentation producing three output:
1 - normal presentation with appendix
2 - presentation without appendix for sharing
3 - handout notes

The .tex file must be in the form of the template downloaded 
running gotalk prepare -p <name>.

Run with the -t flag if your presentation has a name different from that of the folder.
`,
	Run: func(cmd *cobra.Command, args []string) {
		Compile(texName)
	},
}


var (
	publishFolder string
)

var PublishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish the sharable version of the presentation.",
	Long:  `Move the sharable version of the presentation 
	(no hand notes and no appendix) to the PresentationPublic folder
	and update the index.html file in that folder.
	If you want you can change de dafault folder with the -f flag.
	(not available at the moment)
`,
	Run: func(cmd *cobra.Command, args []string) {
		Publish(publishFolder)
	},
}

func InitCommands() () {

	GoTalkCmd.AddCommand(VersionCmd)
	
	GoTalkCmd.AddCommand(PrepareCmd)
	PrepareCmd.Flags().StringVarP(&presentationName, "presentation", "p", "", "Name of the presentation.")
	
	GoTalkCmd.AddCommand(CompileCmd)
	CompileCmd.Flags().StringVarP(&texName, "texName", "t", "", "Name of the .tex file.")
	
	GoTalkCmd.AddCommand(PublishCmd)
	PublishCmd.Flags().StringVarP(&publishFolder, "folder", "f", "", "Folder for the public version of the presentations.")
	
	
}
