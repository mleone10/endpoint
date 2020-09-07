import React from "react";
import DeleteButton from "./DeleteButton";

class ApiKeysList extends React.Component {
  render() {
    if (this.props.apiKeys !== undefined && this.props.apiKeys.length > 0) {
      const keys = this.props.apiKeys.map((key) => (
        <li key={key.key}>
          <p>
            {key.nickname} - {key.readOnly ? "true" : "false"} - {key.key}
          </p>
          <DeleteButton
            keyValue={key.key}
            onDeleteApiKey={this.props.onDeleteApiKey}
          />
        </li>
      ));
      return <ul>{keys}</ul>;
    } else {
      return <p>no api keys found</p>;
    }
  }
}

export default ApiKeysList;
