import {
	constructMessage,
	UI_THEME,
	WS_CONNECTION,
} from "../index.js"
import {
	SHA256
} from "../sha256.js"
import {
	inputWithLabel
} from "./inputWithLabel.js"

export class register extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			username: "",
			password: "",
			password_r: "",
			email: "",
			secret_q: "",
			secret_a: "",
		}
	}

	register_click = () => {
		if (this.username_check() && this.password_check() && this.password_r_check() && this.email_check() && this.secret_q_check() && this.secret_a_check)
			WS_CONNECTION.send(constructMessage( "register",  this.state.username, SHA256(this.state.password), this.state.email, this.state.secret_q, this.state.secret_a.length > 0 ? SHA256(this.state.secret_a) : "") )
	}
	username_check = () => {
		return (this.state.username.length > 4 && this.state.username.length <= 20) && /^[A-Za-z0-9]*$/.test(this.state.username)
	}
	password_check = () => {
		return (this.state.password.length > 7 && this.state.password.length <= 30) && /^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]*$/.test(this.state.password)
	}
	password_r_check = () => {
		return this.state.password_r === this.state.password
	}
	email_check = () => {
		return this.state.email.length === 0 || /^[\w._-]+[+]?[\w._-]+@[\w.-]+\.[a-zA-Z]{2,6}$/.test(this.state.email)
	}
	secret_q_check = () => {
		return this.state.secret_q.length < 50 && (this.state.secret_q.length === 0 || (this.state.secret_q.length > 0 && this.state.secret_a.length > 0))
	}
	secret_a_check = () => {
		return this.state.secret_a.length < 50 && (this.state.secret_a.length === 0 || (this.state.secret_a.length > 0 && this.state.secret_q.length > 0))
	}

	render() {
		return React.createElement("div", {
				onKeyDown: e => {
					if (e.key === "Enter")
						this.register_click()
				},
				style: {
					gap : "1vmin",
					borderRadius: ".4vmin",
					border: "solid " + (this.props.loginData.uiTheme === UI_THEME.LIGHT ? "black " : "white ") + "1px",
					alignItems: "center",
					display: "flex",
					padding: "1vmin",
					flexDirection: "column",
					backgroundColor: this.props.loginData.uiTheme === UI_THEME.LIGHT ? "gainsboro" : "#555555",
					marginTop : "5vmin"
				}
			},
			React.createElement("b", {
					style: {
						alignItems: "center",
						display: "flex",
						fontSize : "2vmin",
						flexDirection: "row",
						marginBottom: "1vmin",
					}
				}, 
				"register"
			),
			React.createElement(inputWithLabel, {
					label : this.state.username.length === 0 || this.username_check() ? "" : "5-20 characters, a-z,0-9",
					changeEvent : (input) => this.setState({username: input}),
					value : this.state.username,
					placeholder : "username"
				}
			),
			React.createElement(inputWithLabel, {
					label : this.state.password.length === 0 || this.password_check() ? "" : "8-30 characters, a-z, 0-9, " + '!@#$%^&*()_+\-=\[\]{};' + "'" + ':"\\|,.<>\/?]*$',
					changeEvent : (input) => this.setState({password: input}),
					value : this.state.password,
					placeholder : "password",
					password : true,
				}
			),
			React.createElement(inputWithLabel, {
					label : this.password_r_check() ? "" : "password and repeat password must match",
					changeEvent : (input) => this.setState({password_r: input}),
					value : this.state.password_r,
					placeholder : "repeat password",
					password : true,
				}
			),
			React.createElement(inputWithLabel, {
					label : this.state.email.length === 0 || this.email_check() ? "" : "incorrect email format",
					changeEvent : (input) => this.setState({email: input}),
					value : this.state.email,
					placeholder : "(optional) email"
				}
			),
			React.createElement(inputWithLabel, {
					label : this.secret_q_check() ? "" : "up to 50 characters and secret answer must by set",
					changeEvent : (input) => this.setState({secret_q: input}),
					value : this.state.secret_q,
					placeholder : "(optional) secret question"
				}
			),
			React.createElement(inputWithLabel, {
					label : this.secret_a_check() ? "" : "up to 50 characters  and secret question must by set",
					changeEvent : (input) => this.setState({secret_a: input}),
					value : this.state.secret_a,
					placeholder : "(optional) secret answer"
				}
			),
			React.createElement("div", {
					style: {
						display: "flex",
						alignItems: "center",
						justifyContent: "center",
						marginTop: ".5vmin",
						flexDirection: "column"
					}
				},
				React.createElement("input", {
						disabled: (!this.username_check() || !this.password_check() || !this.password_r_check() || !this.email_check() || !this.secret_q_check() || !this.secret_a_check()),
						onClick: e => this.register_click(),
						style: {
							display: "flex",
							fontSize: "3vmin",
						},
						type: "button",
						value: "register",
					}, 
				),
			),
		)
	}
}