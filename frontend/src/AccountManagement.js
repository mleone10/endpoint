import React from "react";

function AccountManagement(props) {
  return (
    <div>
      <p>account management</p>
      {props.idToken !== undefined && (<p>ID Token: {props.idToken}</p>)}
    </div>
  );
}

export default AccountManagement;
