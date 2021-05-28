# Welcome to sns2ps - Create Practiscore competitor registration lists from Shoot 'n Score It

## What does this do?

sns2ps (Shoot 'n Score It to Practiscore) allows you to create CSV files from Shoot 'n Score It that can be imported into Practiscore when creating matches.

It supports the following match types:

* Action Air
* Minirifle
* Shotgun

Support for handgun matches is coming soon.

It only supports single-firearm matches.

## Where do I get it?

Visit the [releases page](https://github.com/qawsgh/sns2ps/releases), find the latest release (newest is higher up), and download the appropriate version for your operating system. In case you're not sure, the options are:

* darwin_arm64 - Fancy new Apple computers with Apple's own chip (2021 onwards)
* darwin_amd64 - Older Apple laptops with Intel chips
* linux_amd64 - Linux
* windows_amd64 - Windows computers

## How do I use it?

When you download the release for your computer, you'll get a zip file. Unzip this with the tool of your choice, normally by double clicking it or running `unzip <downloadedfilename>.zip`

There should be a single file in the zip named `sns2ps` for Apple or Linux computers, or `sns2ps.exe` for Windows users. Put this file somewhere you can find it later, like your desktop.

### Linux and Mac only - make the tool runnable

* Open a terminal application
  * For Mac users, you'll want something like [iTerm](https://iterm2.com/) or the `Terminal` under _Applications -> Utilities -> Terminal.app_
* Change to the directory where you extracted the `sns2ps` file from the `zip` file you downloaded
  * If you dragged the `sns` file to your desktop, type `cd ~/Desktop` to get to the right place
* Execute the following command to make the file runnable as your user: `chmod 700 sns2ps`

### Running the tool

On Linux or Macs, you'll need to start with a terminal application like you did when making the file runnable. On Windows, you should run the `Command Prompt`. You can do this by pressing the Windows key on your keyboard or clicking the Windows button on the taskbar, then typing `cmd`. Select `Command Prompt` from the results.

Once you've got a terminal or command prompt running, you must change to the folder where you put the `sns2ps` file after you unzipped it. If you dragged this to your desktop on Windows, you would type `cd Desktop`

To run the tool, type `./sns2ps` on Linux and Mac machines, or `sns2ps.exe` on Windows and press enter.

### Feeding in your match details

If you just followed the instructions above, you will now be prompted for a number of things including your Match ID, your Shoot 'n Score It username, and your password for Shoot 'n Score It.

You can find your MatchID by visiting the main page of your match and looking at your URL. The URL in your browser will be something like `https://shootnscoreit.com/event/22/19991/`

The number after _/event/22/_ is your matchID - in the example above, it would be *19991*

If you enter all of your details correctly, you'll see something like

```
Generating competitor list for "My Awesome Match"
Found X squads
Found Y competitors

Creating CSV named "My_Awesome_Match.csv"
Finished creating competitor csv - you can now import this to Practiscore
```

You'll now have a CSV in the same folder where you ran the command from that you can import into Practiscore. The file will be the name of your match with `_` in place of spaces.

#### These questions are getting tedious - do I REALLY have to answer them every time?

Glad you asked - no, you do not. If you run `sns2ps -v` or `sns2ps.exe -h` you'll see some help. You can supply matchID, username and password on the commandline. If you only provide _some_ of these, you'll be prompted for the rest.

You may wish to do this to protect your Shoot 'n Score It password from being discoverable later.

## Help - I'm stuck

I'm sorry to hear that - I really want this to be as simple to use as possible. If you're stuck, please [raise an issue](https://github.com/qawsgh/sns2ps/issues) here with as much detail as possible.
