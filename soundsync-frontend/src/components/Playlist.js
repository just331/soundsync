import React from 'react'
import AddSong from './AddSong'

export default class App extends React.Component{
  constructor(props) {
    super(props);
    this.state = {  
      songs: {
        'fruit-1' : 'orange',
        'fruit-2' : 'apple'
      }
    }
  }

    queueSong = (song) => {
     //create a unike key for each new fruit item
     var timestamp = (new Date()).getTime();
     // update the state object
     this.state.songs['song-' + timestamp ] = song;
     // set the state
     this.setState({ songs : this.state.songs });
    }
    render() {
      return (
        <div className="component-wrapper">
          <Playlist songs={this.state.songs} />
          <AddSong queueSong={this.queueSong} />
        </div>
      );
     }
    };

    class Playlist extends React.Component{
      render() {
        return (
          <div className="container">
            <ul className="list-group text-center">
              {
                Object.keys(this.props.songs).map(function(key) {
                  return <li className="list-group-item list-group-item-info">{this.props.songs[key]}</li>
                }.bind(this))
              }
            </ul>
           </div>
         );
       }
     }