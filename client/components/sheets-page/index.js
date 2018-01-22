import React, { Component } from 'react';
import Helmet from 'react-helmet';
import { Link } from 'react-router';
import { example, p, link } from '../homepage/styles';

export default class SheetsPage extends Component {
  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    // Load here any data.
    callback(); // this call is important, don't forget it
  }
  /*eslint-enable */

  constructor(props) {
      super(props);

      this.state = {
        googleAuthLink: "",
        isLoading: true,
        error: null,
      };
    }

  componentDidMount() {
      this.setState({ isLoading: true });

      fetch('http://localhost:5001/api/v1/auth')      .then(response => {
        if (response.ok) {
          return response.text();
        } else {
          throw new Error('Something went wrong ...');
        }
      })
.then(data => this.setState({ googleAuthLink: data, isLoading: false }))
            .catch(error => this.setState({ error, isLoading: false }));
  }
  render() {
      const { googleAuthLink, isLoading, error } = this.state;

    if (error) {
      return <p>{error.message}</p>;
    }
    if (~isLoading) {
    return <div>
      <Helmet
        title="Here's all your sheets"
        meta={[
          {
            property: 'og:title',
            content: 'Real life spreadsheets'
          }
        ]} />
      <br />

      <a href={googleAuthLink} className={link} ><button>Login with Google!</button></a>
      <p className={p}>
        Please take a look at <Link className={link} to='/docs'>usage</Link> page.
      </p>
    </div>;
} else {
      return <div><p className={p}>
        Please take a look at <Link className={link} to='/docs'>usage</Link> page.
      </p>
    </div>;
}
}
  }
