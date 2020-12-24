import React from "react";

class ApiKeysList extends React.Component {
  render() {
    if (this.props.apiKeys !== undefined && this.props.apiKeys.length > 0) {
      const keys = this.props.apiKeys.map((key) => (
        <li key={key.key} className="apiKeyItem">
          <p className="apiKey">{key.key}</p>
          {key.readOnly && <p className="readOnly">Read Only</p>}
          <DeleteButton
            keyValue={key.key}
            onDeleteApiKey={this.props.onDeleteApiKey}
          />
        </li>
      ));
      return (
        <div>
          <ul className="apiKeysList">{keys}</ul>
          {this.props.apiKeys.length > 0 && (
            <DeleteKeysButton onDeleteApiKeys={this.props.onDeleteApiKeys} />
          )}
        </div>
      );
    } else {
      return <h3 className="noKeysFound">No API keys found</h3>;
    }
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
      <form className="deleteButton" onSubmit={this.handleSubmit}>
        <button type="submit">Delete</button>
      </form>
    );
  }
}

function DeleteKeysButton(props) {
  return (
    <form className="deleteKeysForm" onSubmit={props.onDeleteApiKeys}>
      <button className="deleteKeysButton" type="submit">
        Delete All Keys
      </button>
    </form>
  );
}

export default ApiKeysList;
