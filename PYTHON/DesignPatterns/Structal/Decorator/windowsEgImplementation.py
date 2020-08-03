from windowsPyHelpers import Window, VerticalSBDecorator, HorizontalSBDecorator, BorderDecorator

w = Window()

wh = HorizontalSBDecorator(w)

whv = VerticalSBDecorator(wh)

whvb = BorderDecorator(whv)

whvb.build()

# or
print()
print('or')
print()

window = BorderDecorator(
            VerticalSBDecorator(
                HorizontalSBDecorator(
                    Window()
        )
    )
)

window.build()