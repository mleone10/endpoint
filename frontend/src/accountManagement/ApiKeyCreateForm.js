import React from "react";

class ApiKeyCreateForm extends React.Component {
  initialState = { readOnly: false };

  constructor(props) {
    super(props);
    this.state = this.initialState;

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({ readOnly: event.target.checked });
  }

  handleSubmit(event) {
    this.props.onCreateNewApiKey(this.state.readOnly);
    this.setState(this.initialState);
    event.preventDefault();
  }

  render() {
    return (
      <div className="apiKeyCreateForm">
        <h3 className="formTitle">Create a new API Key</h3>
        <form onSubmit={this.handleSubmit}>
          <label>
            Read Only:
            <input
              type="checkbox"
              checked={this.state.readOnly}
              onChange={this.handleChange}
              name="readOnly"
            />
          </label>
          <input className="submit" type="submit" />
        </form>
      </div>
    );
  }
}

export default ApiKeyCreateForm;
