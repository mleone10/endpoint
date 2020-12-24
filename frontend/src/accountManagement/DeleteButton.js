import React from "react";

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

export default DeleteButton;
