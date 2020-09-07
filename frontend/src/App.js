import React from "react";
import { HashRouter as Router, Route, Link } from "react-router-dom";
import AuthButton from "./AuthButton";
import About from "./About";
import Documentation from "./Documentation";
import AccountManagement from "./accountManagement/AccountManagement";

class App extends React.Component {
  state = {};

  handleLogin = (idToken) => {
    this.setState({ idToken: idToken });
  };

  handleLogout = () => {
    this.setState({ idToken: undefined });
  };

  render() {
    return (
      <Router>
        <Header onLogin={this.handleLogin} onLogout={this.handleLogout} />
        <NavBar />
        <Content idToken={this.state.idToken} />
      </Router>
    );
  }
}

function Header(props) {
  return (
    <div>
      <AuthButton onLogin={props.onLogin} onLogout={props.onLogout} />
      <Link to="/">
        <h1>endpoint://</h1>
      </Link>
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
  );
}

function Content(props) {
  return (
    <div>
      <Route exact={true} path="/"></Route>
      <Route exact={true} path="/about">
        <About />
      </Route>
      <Route exact={true} path="/docs">
        <Documentation />
      </Route>
      <Route exact={true} path="/acct">
        <AccountManagement idToken={props.idToken} />
      </Route>
    </div>
  );
}

export default App;
