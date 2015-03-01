goair
==========

RTSP Session

After restarting iTunes, a typical Airtunes RTSP session looks like this:

User (not playing music) clicks on AirTunes device.
iTunes requests the OPTIONS method, loaded with an Apple-Challenge header. If this is answered incorrectly, iTunes will display an error message to the user and not enable AirTunes. If this is answered correctly, the AirTunes blue active symbol is active. In both cases, iTunes closes the TCP connection afterwards.
The user starts to play some music. An ANNOUNCE is sent, about the audio and encryption parameters. Then a SETUP call to arrange ports. A RECORD call then signals that the audio session is starting. After that, two SET_PARAMETER calls, for volume (the same volume).
If the user does nothing now, an OPTIONS call is made every 15 seconds, without any Apple-Challenge header.
If the user skips a song, a FLUSH call is made.


Airplay server in Go.
