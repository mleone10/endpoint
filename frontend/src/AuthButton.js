import React from "react";
import firebase from "firebase/app";
import "firebase/auth";
import StyledFirebaseAuth from "react-firebaseui/StyledFirebaseAuth";

const firebaseConfig = {
  apiKey: "AIzaSyBbh8u2Bki1NgjbI2q0wMW2BQZEuzduKHU",
  authDomain: "endpoint-288302.firebaseapp.com",
  databaseURL: "https://endpoint-288302.firebaseio.com",
  projectId: "endpoint-288302",
  storageBucket: "endpoint-288302.appspot.com",
  messagingSenderId: "1059061296434",
  appId: "1:1059061296434:web:db20d46ab09dc50417fb2b",
};
const firebaseApp = firebase.initializeApp(firebaseConfig);

class AuthButton extends React.Component {
  uiConfig = {
    signInOptions: [firebase.auth.GoogleAuthProvider.PROVIDER_ID],
    callbacks: {
      signInSuccessWithAuthResult: () => false,
    },
  };

  componentDidMount() {
    this.unregisterAuthObserver = firebaseApp
      .auth()
      .onAuthStateChanged((user) => {
        this.setState({ isSignedIn: !!user });
        if (user) {
          firebase.auth().currentUser.getIdToken().then(this.props.onLogin)
        } else {
          this.props.onLogout()
        }
      });
  }

  componentWillUnmount() {
    this.unregisterAuthObserver();
  }

  state = {
    isSignedIn: undefined,
  };

  render() {
    return (
      <div>
        {this.state.isSignedIn !== undefined && !this.state.isSignedIn && (
          <div>
            <StyledFirebaseAuth
              uiConfig={this.uiConfig}
              firebaseAuth={firebase.auth()}
            />
          </div>
        )}
        {this.state.isSignedIn && (
          <div>
            <button onClick={() => firebaseApp.auth().signOut()}>
              Sign-out
            </button>
          </div>
        )}
      </div>
    );
  }
}

export default AuthButton;
