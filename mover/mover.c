#include "mover.h"

void move_window(AXUIElementRef window, int x, int y)
{
  // Get the current position
  CGPoint position;
  AXValueRef positionRef;
  AXUIElementCopyAttributeValue(window, kAXPositionAttribute, (CFTypeRef *)&positionRef);
  AXValueGetValue(positionRef, kAXValueCGPointType, &position);

  // Modify the position
  position.x += x;
  position.y += y;

  // Set the new position
  AXValueRef newPositionRef = AXValueCreate(kAXValueCGPointType, &position);
  AXUIElementSetAttributeValue(window, kAXPositionAttribute, newPositionRef);

  CFRelease(positionRef);
  CFRelease(newPositionRef);
}