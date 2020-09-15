import React from "react";
import { HashRouter as Router, Route, Link } from "react-router-dom";

// import AuthButton from "./AuthButton";
import Home from "./Home";
import About from "./About";
import Documentation from "./Documentation";
import AccountManagement from "./accountManagement/AccountManagement";

import "./App.css";

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
        <TopBar onLogin={this.handleLogin} onLogout={this.handleLogout} />
        <NavBar />
        <Content idToken={this.state.idToken} />
        <Footer />
      </Router>
    );
  }
}

function TopBar(props) {
  return (
    <div className="topBar">
      <h1 className="title">
        <Link to="/">endpoint://</Link>
      </h1>
      {/* <AuthButton onLogin={props.onLogin} onLogout={props.onLogout} /> */}
    </div>
  );
}

function NavBar(props) {
  return (
    <div className="navBar">
      <NavBarItem display="about" path="/about" addClass="navBarItemFirst" />
      <NavBarItem display="docs" path="/docs" />
      <NavBarItem display="acct" path="/acct" addClass="navBarItemLast" />
    </div>
  );
}

function NavBarItem(props) {
  return (
    <div className={`navBarItem ${props.addClass ? props.addClass : ""}`}>
      <Link to={props.path}>{props.display}</Link>
    </div>
  );
}

function Content(props) {
  return (
    <div>
      <Route exact={true} path="/">
        <Home />
      </Route>
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

function Footer(props) {
  return (
    <footer>
      Copyright &copy; 2020 Mario Leone &middot; Hosted on{" "}
      <a href="https://github.com/mleone10/endpoint">GitHub</a>
    </footer>
  );
}

export default App;
