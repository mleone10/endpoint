import React from "react";

function About(props) {
  return (
    <div className="content about">
      <h2>&gt; Welcome, Captain</h2>
      <p>
        As the latest commander of a fully self-sustaining space station, all
        eyes are on you to ensure the safety and wellbeing of its inhabitants.
        Life support, defense, industry, and residential concerns fall to you.
      </p>
      <p>One small problem...</p>
      <p>
        The nature of interstellar travel (and exorbitant insurance premiums)
        means we can't possibly risk sending our best and brightest captains to
        the stations themselves. That's where the <b>Endpoint API</b> comes in.
      </p>
      <h3>
        &gt; Introducing The Endpoint&reg; Astrohabitation Platform Interface
      </h3>
      <p>
        The <b>Endpoint API</b> is a state of the art interface for all your
        station management needs. Built on the industry-standard Hypertext
        Transfer Protocol (HTTP), the <b>Endpoint API</b> exposes all facets of
        the station via a set of powerful and intuitive paths.
      </p>
      <h3>&gt; Fully Customizable!</h3>
      <p>
        Of course, your station is bound for greatness - eventually you'll
        outgrow the humble capabilities of the <b>Endpoint API</b>! Fear not,
        for the near-universal adoption of HTTP interfaces means you can write
        your own custom tooling!
      </p>
      <p>
        Try your hand at a command line tool in your favorite language, or a web
        app using the latest framework. Engineer a mobile application to take
        your station management on the go, or automate yourself out of a job
        with a pure server-side solution.
      </p>
      <h3>The station is yours, Captain. Good luck!</h3>
    </div>
  );
}

export default About;
