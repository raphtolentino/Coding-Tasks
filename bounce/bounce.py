import sys, pygame

pygame.init()

size = width, height = 640, 480
dx = 0.2
dy = 0.2
x= 163
y = 120
black = (1,80,75)
white = (255,255,255)

screen = pygame.display.set_mode(size)

while 1:

    for event in pygame.event.get():
        if event.type == pygame.QUIT: sys.exit()

    x += dx
    y += dy

    if x < 0 or x > width:   
        dx = -dx

    if y < 0 or y > height:
        dy = -dy

    screen.fill(black)

    pygame.draw.circle(screen, white, (x,y), 8)

    pygame.display.flip()