# Santa 🎅🏼

A cli to download aoc inputs for those in a hurry or just lazy.

### Installation

You can install the appropriate binary from the [releases page](https://github.com/somnek/santa/releases).

#### Note:

If you're on macOS, you may need to run `xattr -c ./santa_Darwin_x86_64.tar.gz` to (to avoid "unknown developer" warning)

### 🚀 Get started:

1. Retrieve your session token from AoC's cookies:

   - Open AoC website.
   - Right-click, select "Inspect Element."
   - Navigate to the "Application" tab.
   - Find and copy the value of the "session" cookie.

2. Set the session using the following command:

   ```bash
   santa session <session>
   ```

3. download input for specific day & pipe it to output file:

   ```bash
   santa day <day> > <file>
   ```

### Example:

To set the session:

```bash
santa session abcdefghijklmnopqrstuvwxyz0123456789
```

To download input for day 1 & save it to `input.txt`:

```bash
santa day 1 > input.txt
```

