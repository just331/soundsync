import React, { Component } from "react";
import Logo from './logo.png';



class Input extends Component {
  constructor() {
        super(props);
        this.state = {
            name: this.props.name,
            open: false,
            explicitswitch: true,
            style : {width: 0}
    };
    }

    render() {
        return (
          <input
          name="Name"
          type="tel"
          placeholder="Enter Phone Number"
          required="required"
          value={this.state.numberOfGuests}
          onChange={this.handleInputChange} />
    );
  }
}

class Create extends Component {
  constructor() {
        super();
        this.state = {
            customers: [],
            open: false,
            explicitswitch: true,
            style : {width: 0}
    };
    }

    render() {
        return (
        <div>
        <img src={Logo} alt="soundsync Logo" class="logo"></img>
            <h1>soundsync</h1>
                <form method="post">
                    <Input name="Josh">
                    <input
                    name="Name"
                    type="string"
                    placeholder="Enter Username"
                    required="required"
                    value={this.state.numberOfGuests}
                    onChange={this.handleInputChange} />
                    <input
                    name="Name"
                    type="string"
                    placeholder="Enter Party Name"
                    required="required"
                    value={this.state.numberOfGuests}
                    onChange={this.handleInputChange} />
                    <br></br>
                    <input type="submit" value="Send Code"></input>
                </form>
      </div>
    );
  }
}

export default Create;
