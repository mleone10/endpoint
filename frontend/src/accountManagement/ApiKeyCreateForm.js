import React from "react";

class ApiKeyCreateForm extends React.Component {
  initialState = { nickname: "", readOnly: false };

  constructor(props) {
    super(props);
    this.state = this.initialState;

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({
      [event.target.name]:
        event.target.type === "checkbox"
          ? event.target.checked
          : event.target.value,
    });
  }

  handleSubmit(event) {
    this.props.onCreateNewApiKey(this.state.nickname, this.state.readOnly);
    this.setState(this.initialState);
    event.preventDefault();
  }

  render() {
    return (
      <div>
        <p>create new api key</p>
        <form onSubmit={this.handleSubmit}>
          <label>
            nickname:
            <input
              type="text"
              value={this.state.nickname}
              onChange={this.handleChange}
              name="nickname"
            />
          </label>
          <label>
            read only:
            <input
              type="checkbox"
              checked={this.state.readOnly}
              onChange={this.handleChange}
              name="readOnly"
            />
          </label>
          <input type="submit" />
        </form>
      </div>
    );
  }
}

export default ApiKeyCreateForm;
