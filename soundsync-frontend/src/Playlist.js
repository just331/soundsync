import React, { Component } from "react";
import Logo from './logo.png';

class Playlist extends Component {
  render() {
    return (
      <div>
        <ul class="header">
          <li><img src={Logo} alt="soundsync Logo" class="logoicon"></img></li>
          <li>soundsync</li>
          <li>Code: 4DR2</li>
        </ul>
        <ul class="menu">
          <li><button><i class="fa fa-arrow-left"></i></button></li>
          <li>Adriana's Playlist</li>
        </ul>
        <div>
          
        </div>
        <h2>GOT QUESTIONS?</h2>
        <p>The easiest thing to do is post on
        our <a href="http://forum.kirupa.com">forums</a>.
        </p>
      </div>
    );
  }
}
 
export default Playlist;