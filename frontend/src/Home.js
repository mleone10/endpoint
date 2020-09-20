import React from "react";

function Home(props) {
  return (
    <div className="content">
      <h2>&gt; What is Endpoint?</h2>
      <p>
        <b>Endpoint</b> is a space station management simulator.
      </p>
      <p>
        <b>Endpoint</b> is an open API.
      </p>
      <p>
        <b>Endpoint</b> is a blank canvas, a platform ready for you to build
        upon.
      </p>
      <p>
        At its heart, <b>Endpoint</b> is an open source REST API written in Go
        which simulates the various systems, inhabitants, and challenges
        associated with a self-sustaining space station in the distant future.
        You are the Captain, tasked with the stations's success. Will you mine
        the local asteroid belt for resources? Will you become a manufacturing
        hub? Will you build your arsenal and take what you need by force? The
        decision is yours.
      </p>
      <p>
        Most importantly, <b>Endpoint</b> leaves the creation of a user
        interface in the hands of its community. Fledgling developers can use{" "}
        <b>Endpoint</b> as an underlying system as they make their first foray
        into programming. Experienced engineers can use it as a backend service
        as they learn new frameworks or technologies.
      </p>
      <p>Check out the docs to get started.</p>
    </div>
  );
}

export default Home;
