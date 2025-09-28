// SETTINGS

#set text(
  font: "Cormorant Garamond",
  lang: "es",
  size: 12pt,
  weight: "regular"

)

#set par(
  justify: true
)

#show math.equation: set text(weight: "extralight")

#show heading.where(level: 1): set text(font: "Cinzel Decorative", size: 18pt, weight: "regular")
#show heading.where(
  level: 2
): it => block(width: 100%)[
  #set text(15pt, weight: "bold")
  #smallcaps(it.body)
]
#show heading.where(
  level: 3
): it => block(width: 100%)[
  #set text(12pt, weight: "bold")
  #smallcaps(it.body)
]
#show heading.where(
  level: 4
): it => block(width: 100%)[
  #set text(12pt, weight: "bold")
  #text(it.body)
]

#set table(
  stroke: (thickness: 0.5pt)
)

// CONSTANTS

#let title = "Template"
#let subject = "Subject Name"
#let author = "Alejandro Benítez Bravo \n A01712835"
#let professor = "Name of the teacher"
#let date = datetime.today()

// HEADER AND FOOTER SETUP

#set page(
  header: context { if(counter(page).get().at(0)== 1) [
      #set align(center)
    ] else [
    #{
      set par(spacing: 0.2cm)
       subject
       h(1fr)
       title
       line(length: 100%, stroke: 0.3pt)
    }
    ]
  }, 
  footer: context { if(counter(page).get().at(0)== 1) [
      #set align(center)
    ] else [
      #box(image("img/tec-logo.png", height: 40%)) #h(0.3cm) Cámpus Queŕetaro #h(1fr) #text(size: 16pt)[#counter(page).display()]
    ]
  }, 
)

// TITLE

#{
   set par(spacing: 0.7cm)
   align(center + horizon)[
    #image("img/tec-logo.png", width: 30%)
    #line(length: 100%, stroke: 0.3pt)
    #text(size: 32pt, font: "Cinzel Decorative", weight: "extralight")[#title]
    #line(length: 100%, stroke: 0.3pt)

    #upper(text(size: 18pt)[#subject])

    #text(author, size: 14pt)

    #text(professor, size: 14pt)

    #date.display("[month repr:long] [day], [year]")
  ]
  pagebreak()
}

= Introduction

#lorem(100)

= Body

#lorem(50)

== Subsection 1

#lorem(100)

=== Sub-sub section 1

==== Crazy

#figure(
  table(
    columns: (auto, auto, auto, auto),
    table.header([One], [two], [three], [four]),
  ),
  caption: [Hello]
)

#math.equation($ x - y times x $, block: true, numbering: "(1)")

#math.equation($ x - y times x $, block: true, numbering: "(1)")
