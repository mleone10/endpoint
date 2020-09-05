import React from "react";
import AuthButton from "./AuthButton";

class App extends React.Component {
  handleLogin = (idToken) => {
    this.setState({ idToken: idToken });
    fetch("https://api.endpointgame.com/user/api-keys", {
      headers: { Authorization: idToken },
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
      });
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
