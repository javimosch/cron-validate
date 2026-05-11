package main
import ("fmt";"os";"strconv";"strings")
func main() {
	if len(os.Args) < 2 { fmt.Fprintln(os.Stderr,"Usage: cron-validate <cron_expr>"); os.Exit(1) }
	expr := os.Args[1]
	parts := strings.Fields(expr)
	if len(parts) != 5 { fmt.Fprintln(os.Stderr,"INVALID — expected 5 fields, got",len(parts)); os.Exit(1) }
	names := []string{"minute","hour","day of month","month","day of week"}
	for i, p := range parts {
		if p == "*" { continue }
		if _, err := strconv.Atoi(p); err != nil {
			if p != "*" && !strings.ContainsAny(p, ",-/") {
				fmt.Fprintf(os.Stderr,"INVALID — field %d (%s): %s\n", i+1, names[i], p)
				os.Exit(1)
			}
		}
	}
	fmt.Println("VALID")
	fmt.Printf("  Minute: %s\n", parts[0])
	fmt.Printf("  Hour: %s\n", parts[1])
	fmt.Printf("  Day of Month: %s\n", parts[2])
	fmt.Printf("  Month: %s\n", parts[3])
	fmt.Printf("  Day of Week: %s\n", parts[4])
}
