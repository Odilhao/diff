package diff

// diff func returns a slice with the elements whose are in slice a but not in slice b
func DiffSet(a []string, b []string) (diff []string) {

  for i := 0; i < len(a); i++ {
    found := false
    for j := 0; j < len(b); j++ {
      if a[i] == b[j] {
        found = true
        break
      }
    }
    if !found {
      diff = append(diff, a[i])
    }
  }
  return
}
