import React from "react";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTwitter } from "@fortawesome/free-brands-svg-icons";

class Footer extends React.Component {
  render() {
    return (
      <footer>
        Copyright &copy; 2020 <a href="https://marioleone.me">Mario Leone</a>{" "}
        &middot;{" "}
        <a href="https://twitter.com/endpointgame">
          <FontAwesomeIcon className="icon" icon={faTwitter} />
        </a>{" "}
        &middot; Hosted on{" "}
        <a href="https://github.com/mleone10/endpoint">GitHub</a>
      </footer>
    );
  }
}

export default Footer;
