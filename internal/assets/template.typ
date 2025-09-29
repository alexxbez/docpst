#let essay(
  title: none,
  authors: (),
  doc,
) = {
  // Simple page setup
  set page(margin: 1in)
  set text(size: 12pt)
  
  // Title and author
  set align(center)
  text(16pt, weight: "bold", title)
  v(12pt)
  
  for author in authors {
    text(author.name)
    v(6pt)
  }
  
  v(24pt)
  
  // Main content
  set par(justify: true)
  set align(left)
  doc
}
