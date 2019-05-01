import React, { useState } from 'react'
import { withStyles } from '@material-ui/core/styles'
import SoundSyncInput from 'components/Input'
import Grid from '@material-ui/core/Grid'

const styles = (theme) => {
  return {
    Container: {
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      padding: '135px 0px',
    },
    dense: {
      marginTop: 19,
    },
    menu: {
      width: 200,
    },
  }
}

// searchResults Type
//
//  {
//    Tracks : {
//      Items: [
//        {
//          TrackName : string,
//          Artists: [
//            {
//              ArtistID : string,
//              ArtistName : string
//            }, {...}, {...}
//          ],
//          Explicit : bool,
//          TrackID : string,
//          TrackURI : string
//        }, {...}, {...}
//      ]
//    }
//  }
//

function Search({ classes }) {
  const [values, setValues] = useState({
    searchText: '',
    searchResults: [],
  })

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value })
  }

  const searchSongs = (event) => {
    // make request to api
    // Request URL:  soundsync.xyz/SearchSpotify/{searchText}
    // searches by song title
    // response from api will be searchResults Type above
    // mocked results
    values.searchResults = {
      Tracks: {
        Items: [
          {
            TrackName: 'olaf',
            Artists: [
              {
                ArtistID: '123',
                ArtistName: 'Elsa',
              },
            ],
            Explicit: true,
            TrackID: '321',
            TrackURI: 'spotURIbaby',
          },
          {
            TrackName: 'sven',
            Artists: [
              {
                ArtistID: '456',
                ArtistName: 'Ana',
              },
            ],
            Explicit: true,
            TrackID: '654',
            TrackURI: 'spotURIbaby',
          },
          {
            TrackName: 'squirt',
            Artists: [
              {
                ArtistID: '789',
                ArtistName: 'sledge',
              },
            ],
            Explicit: true,
            TrackID: '987',
            TrackURI: 'spotURIbaby',
          },
        ],
      },
    }

    console.log('searchResults')
    console.log(values.searchResults)
  }

  const Test = ({ tracks }) => (
    <>
      {tracks.Items.map((track) => (
        <Grid item sm={12}>
          <div className='track' key={track.TrackName}>
            {track.TrackName} | {track.Artists[0].ArtistName}
          </div>
        </Grid>
      ))}
    </>
  )

  const handleKeyPress = (event) => {
    if (event.key == 'Enter') {
      console.log('enter pressed!')
      searchSongs()
    }
  }

  return (
    <Grid
      container
      direction='column'
      justify='flex-start'
      alignItems='stretch'
    >
      <Grid item sm={12}>
        <SoundSyncInput
          id='searchText'
          value={values.searchText}
          placeholder='Search'
          onChange={handleChange('searchText')}
          onKeyPress={handleKeyPress}
        />
      </Grid>
      {values.searchResults.length > 0 ? (
        <Test tracks={values.searchResults} />
      ) : (
        <></>
      )}
    </Grid>
  )
}

export default withStyles(styles)(Search)
