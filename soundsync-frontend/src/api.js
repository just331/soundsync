const baseUrl = 'http://localhost:3005'

const SoundSyncAPI = {
  CreateParty: function(nickName, phoneNum, partyName) {
    fetch(
      baseUrl + '/CreateParty/' + nickName + '/' + partyName + '/' + phoneNum,
    )
      .then((res) => res.json())
      .then(
        (result) => {
          return result
        },
        (error) => {
          console.log(error)
        },
      )
  },

  JoinParty: function(nickName, partyCode, phoneNum) {
    fetch(baseUrl + '/JoinParty/' + nickName + '/' + partyCode + '/' + phoneNum)
      .then((res) => res.json())
      .then(
        (result) => {
          if (result === 'Party Joined') {
            console.log('API: Joined Party')
          } else {
            console.log('API: Join Party failed')
          }
        },
        (error) => {
          console.log(error)
        },
      )
  },

  LinkSpotify: function() {
    fetch(baseUrl + '/LinkSpotify').then((res) => {
      return res
    })
  },

  Callback: function() {
    fetch(baseUrl + '/callback').then((res) => {
      return null
    }) // Do something with the result
  },

  Play: function() {
    fetch(baseUrl + '/MediaControls/Play').then((res) => {
      return null
    }) // Do something with the result
  },

  Pause: function() {
    fetch(baseUrl + '/MediaControls/Pause').then((res) => {
      return null
    }) // Do something with the result
  },

  Next: function() {
    fetch(baseUrl + '/MediaControls/Next').then((res) => {
      return null
    }) // Do something with the result
  },

  Previous: function() {
    fetch(baseUrl + '/MediaControls/Previous').then((res) => {
      return null
    }) // Do something with the result
  },

  SearchSpotify: function(query) {
    fetch(baseUrl + '/SearchSpotfy/' + query).then((res) => {
      return null
    }) // Do something with the result
  },

  AddSong: function(songURI) {
    fetch(baseUrl + '/AddSong/' + songURI).then((res) => {
      return null
    }) // Do something with the result
  },
}

export default SoundSyncAPI
