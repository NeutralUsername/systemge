import { 
    UI_THEME 
} from "../index.js"

export class playButton extends React.Component {
	constructor(props) {
		super(props)
	}

	render() {

		return React.createElement("div", {
                style: {
                    gap : ".5vmax",
                    borderRadius: ".2vmax",
                    border: "solid " + (this.props.uiTheme === UI_THEME.LIGHT ? "black " : "white ") + "1px",
                    alignItems: "center",
                    display: "flex",
                    padding: ".5vmax",
                    flexDirection: "column",
                    height : "4vmax",
                    width : "7.5vmax",
                    backgroundColor: this.props.uiTheme === UI_THEME.LIGHT ? "gainsboro" : "#555555",
                }
            },
            React.createElement("input", {
                    type : "button",
                    value : "PLAY\nONLINE",
                    disabled : true,
                    style : {
                        cursor : "pointer",
                        fontSize : "1.3vmax",
                        width : "100%",
                        height : "100%",
                    }
                }, 
            ),
        )
	}
} 


