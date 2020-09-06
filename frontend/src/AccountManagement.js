import React from "react";

class AccountManagement extends React.Component {
  // TODO: Unset data on sign out
  state = {};

  fetchApiKeys = () => {
    // TODO: Create a way to easily point to a local API for development
    if (this.props.idToken !== undefined) {
      fetch("https://api.endpointgame.com/user/api-keys", {
        headers: { Authorization: this.props.idToken },
      })
        .then((res) => res.json())
        .then((data) => {
          this.setState({ apiKeys: data.keys });
        });
    }
    // TODO: Handle failure
  };

  handleCreateNewApiKey = (nickname, readOnly) => {
    fetch("https://api.endpointgame.com/user/api-keys", {
      method: "POST",
      headers: { Authorization: this.props.idToken },
      body: JSON.stringify({
        readOnly: readOnly,
        nickname: nickname,
      }),
    }).then(() => {
      this.fetchApiKeys(this.props.idToken);
    });
    // TODO: Handle failure
  };

  handleDeleteApiKey = (keyValue) => {
    fetch(`https://api.endpointgame.com/user/api-keys/${keyValue}`, {
      method: "DELETE",
      headers: { Authorization: this.props.idToken },
    });
    // TODO: Handle failure
  };

  componentDidMount() {
    if (this.props.idToken !== undefined) {
      this.fetchApiKeys(this.props.idToken);
    }
  }

  componentDidUpdate(prevProps, prevState) {
    if (
      this.props.idToken !== prevProps.idToken ||
      (this.state.apiKeys !== undefined &&
        prevState.apiKeys !== undefined &&
        this.state.apiKeys.length !== prevState.apiKeys.length)
    ) {
      this.fetchApiKeys(this.props.idToken);
    }
  }

  render() {
    if (this.state.apiKeys !== undefined) {
      return (
        <div>
          <ApiKeyPanel
            apiKeys={this.state.apiKeys}
            onDeleteApiKey={this.handleDeleteApiKey}
          />
          <ApiKeyCreatePanel onCreateNewApiKey={this.handleCreateNewApiKey} />
        </div>
      );
    } else {
      return <p>login to manage your account</p>;
    }
  }
}

class ApiKeyPanel extends React.Component {
  render() {
    if (this.props.apiKeys !== undefined && this.props.apiKeys.length > 0) {
      return (
        <ApiKeysList
          apiKeys={this.props.apiKeys}
          onDeleteApiKey={this.props.onDeleteApiKey}
        />
      );
    } else {
      return <p>no api keys found</p>;
    }
  }
}

function ApiKeysList(props) {
  if (props.apiKeys !== undefined) {
    const keys = props.apiKeys.map((key) => (
      <li key={key.key}>
        <p>
          {key.nickname} - {key.readOnly ? "true" : "false"} - {key.key}
        </p>
        <DeleteButton
          keyValue={key.key}
          onDeleteApiKey={props.onDeleteApiKey}
        />
      </li>
    ));
    return <ul>{keys}</ul>;
  }
}

class DeleteButton extends React.Component {
  constructor(props) {
    super(props);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(event) {
    this.props.onDeleteApiKey(this.props.keyValue);
    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <button type="submit">Delete</button>
      </form>
    );
  }
}

class ApiKeyCreatePanel extends React.Component {
  constructor(props) {
    super(props);
    this.state = { nickname: "", readOnly: false };

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
              value={this.state.readOnly}
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

export default AccountManagement;
