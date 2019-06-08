import React from 'react';

class Home extends React.Component {
  render() {
    return (
      <div className="welcome-screen">
        <header className="header">
          <h1>
            Welcome to Agenda!
          </h1>
          <h3>
            Turn-based game of political intrigue played via e-mail.
          </h3>
          <div className="login-buttons">
            <button
              className="button"
              type='button'
              onClick={() => {  this.props.history.push('/new-game') }}>
              Register
            </button>
            <button
              className="button">
              Log in
            </button>
          </div>
        </header>
      </div>
    )
  };
}
export default Home;
