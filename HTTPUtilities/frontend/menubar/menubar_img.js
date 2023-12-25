import {
	WS_CONNECTION,
	UI_THEME
} from "../index.js"

export class menubar_img extends React.Component {
	constructor(props) {
		super(props)
	}

	render() {
		return  React.createElement("img", {
                src: this.props.src,
                onClick: e => this.props.clickEvent(),
                style: {
                    display: "flex",
                    flexDirection: "row",
                    borderRadius : "0.3vmax",
                    alignItems: "center",
                    backgroundColor: this.props.backgroundColor,
                    border: "solid .1vmax" ,
                    padding: "0.1vmax",
                    cursor: "pointer",
                    width: "1.3vmax",
                    height: "1.3vmax",
                }
            },
        )
	}
}

