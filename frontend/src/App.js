import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import AuthButton from "./AuthButton";

class App extends React.Component {
  handleLogin = (idToken) => {
    this.setState({ idToken: idToken });
    // TODO: Create a way to easily point to a local API for development
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
      <Router>
        <Header onLogin={this.handleLogin} onLogout={this.handleLogout} />
        <NavBar />
        <Content />
      </Router>
    );
  }
}

function Header(props) {
  return (
    <div>
      <AuthButton onLogin={props.handleLogin} onLogout={props.handleLogout} />
      <Link to="/"><h1>endpoint://</h1></Link>
    </div>
  );
}

function NavBar(props) {
  return (
    <div>
      <Link to="/about">about</Link>
      <Link to="/docs">docs</Link>
      <Link to="/acct">acct</Link>
    </div>
  )
}

function Content(props) {
  return <div>
    <Route exact={true} path="/">home page</Route>
    <Route exact={true} path="/about">about page</Route>
    <Route exact={true} path="/docs">docs page</Route>
    <Route exact={true} path="/acct">acct page</Route>
  </div>
}

export default App;
