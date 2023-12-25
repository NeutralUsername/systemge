import { 
    UI_THEME, 
    WS_CONNECTION 
} from "../index.js"
import {
    message
} from "./chatMessage.js"

export class chat extends React.Component {
	constructor(props) {
		super(props) 
	}

    scrollToBottom(ref) {
        if (ref) {
            ref.scrollTop = ref.scrollHeight
        }
    }

	render() {
        let messages = []
        if (this.props.communicatorSelectedUser && this.props.pms[this.props.communicatorSelectedUser]) {
            this.props.pms[this.props.communicatorSelectedUser].forEach((m, i) => {
                messages.push(React.createElement(message, {
                    key : i,
                    message : m,
                    fromSelf : m.senderUserId != this.props.communicatorSelectedUser,
                    senderName :  m.senderUserId != this.props.communicatorSelectedUser ? "you" : this.props.subscriptions[this.props.communicatorSelectedUser].user.username, 
                    uiTheme : this.props.uiTheme,
                }))
            })
        }
		return React.createElement("div", {
                ref : (ref) => this.scrollToBottom(ref),

                style  : {
                    overflowWrap : "break-word",
                    wordBreak : "break-all",
                    display : "flex",
                    height : "35vmin",
                    alignItems : "flex-start",
                    flexDirection : "column",
                    outline : "solid 1px",
                    overflow: "scroll",
                    overflowX : "hidden",
                    scrollbarWidth: "none",
                    msOverflowStyle: "none",
                    backgroundColor : this.props.uiTheme == UI_THEME.LIGHT ? "white" : "#2e3033",
                    borderRadius: ".4vmin",
                    whiteSpace : "pre-wrap",
                }
            },
            messages,
        )
	}
}