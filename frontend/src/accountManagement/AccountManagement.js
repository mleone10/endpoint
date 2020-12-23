import React from "react";
import ApiKeysList from "./ApiKeysList";
import ApiKeyCreateForm from "./ApiKeyCreateForm";
import "./accountManagement.css";

class AccountManagement extends React.Component {
  // TODO: Unset data on sign out
  state = {};

  fetchApiKeys = () => {
    // TODO: Create a way to easily point to a local API for development
    if (this.props.idToken !== undefined) {
      fetch(
        `https://api.endpointgame.com/accounts/${this.props.uid}/api-keys`,
        {
          headers: { Authorization: this.props.idToken },
        }
      )
        .then((res) => res.json())
        .then((data) => {
          this.setState({ apiKeys: data.keys });
        });
    }
    // TODO: Handle failure
  };

  handleCreateNewApiKey = (nickname, readOnly) => {
    fetch(`https://api.endpointgame.com/accounts/${this.props.uid}/api-keys`, {
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
    fetch(
      `https://api.endpointgame.com/accounts/${this.props.uid}/api-keys/${keyValue}`,
      {
        method: "DELETE",
        headers: { Authorization: this.props.idToken },
      }
    ).then(() => {
      this.fetchApiKeys(this.props.idToken);
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
          <ApiKeysList
            apiKeys={this.state.apiKeys}
            onDeleteApiKey={this.handleDeleteApiKey}
          />
          <ApiKeyCreateForm onCreateNewApiKey={this.handleCreateNewApiKey} />
        </div>
      );
    } else {
      return (
        <div className="content">
          <h3 className="loginMsg">Log in to manage your account.</h3>
        </div>
      );
    }
  }
}

export default AccountManagement;
