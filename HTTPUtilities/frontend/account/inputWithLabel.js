export class inputWithLabel extends React.Component {
	constructor(props) {
		super(props)
	}

	render() {
		return React.createElement("div", {
                style: {
                    display: "flex",
                    flexDirection: "row",
                }

            },
            React.createElement("input", {
                    type: this.props.password ? "password" : "input",
                    placeholder: this.props.placeholder,
                    value : this.props.value ,
                    onChange : (e) => this.props.changeEvent(e.target.value),
                    style: {
                        display: "flex",
                        fontSize: "2vmin",
                    },
                }, 
            ),
            React.createElement("label", {
                    style: {
                        marginTop: ".5vmin",
                        marginLeft: "1vmin",
                        display: "flex",
                        position: "absolute",
                        marginLeft: "28vmin"
                    }
                }, 
                this.props.label
            )
        )
	}
}