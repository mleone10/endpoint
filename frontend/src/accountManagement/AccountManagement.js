import React from "react";
import ApiKeysList from "./ApiKeysList";
import ApiKeyCreateForm from "./ApiKeyCreateForm";
import "./accountManagement.css";

class AccountManagement extends React.Component {
  state = { apiKeys: [] };

  fetchApiKeys = () => {
    // TODO: Create a way to easily point to a local API for development
    if (this.props.idToken !== undefined) {
      fetch(
        `https://api.endpointgame.com/accounts/${this.props.uid}/api-keys`,
        {
          headers: { Authorization: `Bearer ${this.props.idToken}` },
        }
      )
        .then((res) => res.json())
        .then((data) => {
          this.setState({ apiKeys: data.apiKeys });
        });
    }
    // TODO: Handle failure
  };

  handleCreateNewApiKey = (nickname, readOnly) => {
    fetch(`https://api.endpointgame.com/accounts/${this.props.uid}/api-keys`, {
      method: "POST",
      headers: { Authorization: `Bearer ${this.props.idToken}` },
      body: JSON.stringify({
        readOnly: readOnly,
        nickname: nickname,
      }),
    }).then(() => {
      this.fetchApiKeys();
    });
    // TODO: Handle failure
  };

  handleDeleteApiKey = (keyValue) => {
    fetch(
      `https://api.endpointgame.com/accounts/${this.props.uid}/api-keys/${keyValue}`,
      {
        method: "DELETE",
        headers: { Authorization: `Bearer ${this.props.idToken}` },
      }
    ).then(() => {
      this.fetchApiKeys();
    });
    // TODO: Handle failure
  };

  componentDidUpdate(prevProps, prevState) {
    if (
      this.props.idToken === undefined &&
      prevProps.idToken !== this.props.idToken
    ) {
      // If we just logged out, clear the list of API Keys
      this.setState({ apiKeys: [] });
    } else if (this.props.idToken !== prevProps.idToken) {
      // If we just logged in, fetch the list of API Keys
      this.fetchApiKeys();
    }
  }

  render() {
    if (this.props.idToken !== undefined) {
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
