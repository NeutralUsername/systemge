import {
    COMMUNICATOR_CONTENT,
    UI_THEME,
	WS_CONNECTION,
    constructMessage
} from "../index.js"

export class input extends React.Component {
	constructor(props) {
		super(props)
	}

    inputValue = () => {
        switch (this.props.loginData.communicatorContent) {
            case COMMUNICATOR_CONTENT.CONTACTS:
                return this.props.communicatorContactsInput
            case COMMUNICATOR_CONTENT.BLOCKS:
                return this.props.communicatorBlocksInput
            case COMMUNICATOR_CONTENT.CHATS:
                return this.props.communicatorSelectedUser !== null ? this.props.communicatorChatsInputs[this.props.communicatorSelectedUser] : ""
        }
    }

    inputPlaceholder = () => {
        switch(this.props.loginData.communicatorContent) {
            case COMMUNICATOR_CONTENT.CONTACTS:
                return "username"
            case COMMUNICATOR_CONTENT.BLOCKS:
                return "username"
            case COMMUNICATOR_CONTENT.CHATS:
                return "message"
        }
    }

    inputChangeEvent = (value) => {
        switch(this.props.loginData.communicatorContent) {
            case COMMUNICATOR_CONTENT.CONTACTS:
                this.props.setStateRoot({communicatorContactsInput : value})
                break;
            case COMMUNICATOR_CONTENT.BLOCKS:
                this.props.setStateRoot({communicatorBlocksInput : value})
                break;
            case COMMUNICATOR_CONTENT.CHATS:
                this.props.setStateRoot({communicatorChatsInputs : {
                    ...this.props.communicatorChatsInputs,
                    [this.props.communicatorSelectedUser] : value
                }})
                break;
        }
    }

    submitButtonEvent = () => {
        switch (this.props.loginData.communicatorContent) {
            case COMMUNICATOR_CONTENT.CONTACTS:
                if (this.props.communicatorContactsInput.trim() != "") {
                    WS_CONNECTION.send(constructMessage("addContact", this.props.communicatorContactsInput))
                    this.props.setStateRoot({communicatorContactsInput : ""})
                }
                break;
            case COMMUNICATOR_CONTENT.BLOCKS:
                if (this.props.communicatorBlocksInput.trim() != "") {
                    WS_CONNECTION.send(constructMessage("addBlock", this.props.communicatorBlocksInput))
                    this.props.setStateRoot({communicatorBlocksInput : ""})
                }
                break;
            case COMMUNICATOR_CONTENT.CHATS:
                if (this.props.communicatorSelectedUser !== null) {
                    if (this.props.communicatorChatsInputs[this.props.communicatorSelectedUser].trim() != "") {
                        WS_CONNECTION.send(constructMessage("addPm", this.props.communicatorSelectedUser, this.props.communicatorChatsInputs[this.props.communicatorSelectedUser]))
                        this.props.setStateRoot({communicatorChatsInputs : {
                            ...this.props.communicatorChatsInputs,
                            [this.props.communicatorSelectedUser] : ""
                        }})
                    }
                }
                break;
        }
    }

    submitButtonText = () => {
        if(this.props.communicatorContent === "contacts")
            return "add"
        if(this.props.communicatorContent === "blocks")
            return "block"
        return "send"
    }

	render() {
		return React.createElement("div", {
                style: {
                    display: "flex",
                    outline: "solid " + (this.props.uiTheme === UI_THEME.LIGHT ? "black" : "white") + " 1px",
                    height : "5vmin",
                    gap :".5vmin",
                    alignItems: "center",
                    justifyContent: "center",
                    flexDirection: "row",
                    borderRadius: ".4vmin", 
                    backgroundColor: this.props.uiTheme === UI_THEME.LIGHT  ? "gainsboro" : "#555555",
                }
            },
            React.createElement("input", {
                    type: "text",
                    placeholder: this.inputPlaceholder(),
                    disabled : this.props.communicatorContent === COMMUNICATOR_CONTENT.CHATS && this.props.communicatorSelectedUser === null,
                    value : this.inputValue(),
                    onFocus : (e) => this.props.setStateRoot({lockKeybinds: true}),
                    onBlur : (e) => this.props.setStateRoot({lockKeybinds: false}),
                    onKeyPress : (e) => {
                        if(e.key === "Enter")
                            this.submitButtonEvent()
                    },
                    onChange : (e) => this.inputChangeEvent(e.target.value),
                    style: {
                        borderRadius  : ".4vmin",
                        display: "flex",
                        alignItems: "center",
                        justifyContent: "center",
                        fontSize: "2vmin",
                        width :"33vmin"
                    }
                }, 
            ),
            React.createElement("button", {
                    onClick : (e) => this.submitButtonEvent(),
                    disabled : this.props.communicatorContent === COMMUNICATOR_CONTENT.CHATS && this.props.communicatorSelectedUser === null,
                    style: {
                        display: "flex",
                        alignItems: "center",
                        justifyContent: "center",
                        fontSize: "2vmin",
                        width : "7vmin",
                        cursor: "pointer",
                    }
                }, 
               this.submitButtonText()
            )
        )
	}
}
