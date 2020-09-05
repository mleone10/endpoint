import React from "react";
import AuthButton from "./AuthButton";

class App extends React.Component {
  handleLogin = (idToken) => {
    this.setState({ idToken: idToken });
  };

  handleLogout = () => {
    this.setState({ idToken: undefined });
  };

  render() {
    return (
      <div>
        <h1>Endpoint</h1>
        <AuthButton onLogin={this.handleLogin} onLogout={this.handleLogout} />
      </div>
    );
  }
}

export default App;
