import React from "react";
import { HashRouter as Router, Route, Link } from "react-router-dom";

// import AuthButton from "./AuthButton";
import Footer from "./Footer";
import Home from "./Home";
import About from "./About";
import Documentation from "./Documentation";
import AccountManagement from "./accountManagement/AccountManagement";

import "./App.css";
import AuthButton from "./AuthButton";

class App extends React.Component {
  state = {};

  handleLogin = (uid, idToken) => {
    console.log(uid);
    this.setState({ uid: uid, idToken: idToken });
  };

  handleLogout = () => {
    this.setState({ idToken: undefined });
  };

  render() {
    return (
      <Router>
        <AuthButton onLogin={this.handleLogin} onLogout={this.handleLogout} />
        <TopBar />
        <NavBar />
        <Content uid={this.state.uid} idToken={this.state.idToken} />
        <Footer />
      </Router>
    );
  }
}

function TopBar(props) {
  // TODO: Add button for Login
  return (
    <div className="topBar">
      <h1 className="title">
        <Link to="/">endpoint://</Link>
      </h1>
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
        <AccountManagement uid={props.uid} idToken={props.idToken} />
      </Route>
    </div>
  );
}

export default App;
