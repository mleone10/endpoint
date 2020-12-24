import React from "react";
import DeleteButton from "./DeleteButton";

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
      return <ul className="apiKeysList">{keys}</ul>;
    } else {
      return <h3 className="noKeysFound">No API keys found</h3>;
    }
  }
}

export default ApiKeysList;
