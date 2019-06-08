import React from 'react';
import './CreateGame.css';

const createPlayersForm = (random) => {
  const nations = ['Prussia', 'England', 'Russia', 'Austria-Hungary', 'Turkey', 'Italy', 'France'];
    return nations.map((nation) => {
      return (
        <div className="nation-row">
          <input type="email" placeholder="Inactive when empty"/>
          { random ? <p> - {nation}</p> : '' }
        </div>
    )
  })
}

class CreateGame extends React.Component {

  constructor(props) {
    super(props);
    this.state = { randomNations: false };
    this.random = this.random.bind(this);
  }

  random() {
    const randomNations = this.state.randomNations;
    this.setState({
      randomNations: !randomNations
    });
  }

  render() {
    return (
      <div className="welcome-screen">
        <header className="header">
          <h1>
            Create new game:
          </h1>
          <button
            className="button"
            type='button'
            onClick={this.random}>
            { this.state.randomNations ? 'Random nations' : 'Select nations' }
          </button>
          <div>
            {createPlayersForm(this.state.randomNations)}
          </div>
          <p>
            Remind players about sending orders after <input className="narrow-input" type="number"/> days of inactivity
          </p>
          <p>
            Kick players from the game after <input className="narrow-input" type="number"/> days of inactivity
          </p>
        </header>
      </div>
    );
  }
}

export default CreateGame;
