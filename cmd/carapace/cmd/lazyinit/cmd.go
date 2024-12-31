package lazyinit

import (
	"fmt"
	"strings"
)

func CmdClink(completers []string) string {
	// TODO pathSnippet
	snippet := `local function carapace_completion(command)
  return function(word, word_index, line_state, match_builder)
    local compline = string.sub(line_state:getline(), 1, line_state:getcursor())
    local prog = string.format("env CARAPACE_COMPLINE=%%s carapace %%s  cmd-clink \"\"", string.format("%%q", compline), command)

    local output = io.popen(prog):read("*a")
    for line in string.gmatch(output, '[^\r\n]+') do
      local matches = {} 
      for m in string.gmatch(line, '[^\t]+') do 
        table.insert(matches, m) 
      end 
      match_builder:addmatch({ 
        match = matches[1], 
        description = matches[2] 
      }) 
    end
    return true
  end
end

%v
`
	argmatchers := make([]string, 0, len(completers))
	for _, completer := range completers {
		argmatchers = append(argmatchers,
			fmt.Sprintf(`clink.argmatcher("%v"):addarg({carapace_completion("%v")}):loop(1)`, completer, completer),
			fmt.Sprintf(`clink.argmatcher("%v.exe"):addarg({carapace_completion("%v")}):loop(1)`, completer, completer),
		)
	}
	return fmt.Sprintf(snippet, strings.Join(argmatchers, "\n"))
}
