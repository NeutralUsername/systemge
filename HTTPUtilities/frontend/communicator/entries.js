import { 
    COMMUNICATOR_CONTENT,
    UI_THEME 
} from "../index.js"

export class entries extends React.Component {
	constructor(props) {
		super(props)
	}

	render() {
        let entries = []
		return React.createElement("div", {
                style  : {
                    display : "flex",
                    height : this.props.loginData.communicatorContent === COMMUNICATOR_CONTENT.CHATS ? "25vmin" : "60vmin",
                    alignItems : "flex-start",
                    flexDirection : "column",
                    outline : "solid 1px",
                    overflow: "scroll",
                    scrollbarWidth: "none",
                    msOverflowStyle: "none",
                    backgroundColor : this.props.loginData.uiTheme === UI_THEME.LIGHT ? "white" : "#2e3033",
                    borderRadius: ".4vmin",
                }
            },
            entries
        )
	}
}