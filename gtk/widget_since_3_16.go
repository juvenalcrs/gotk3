// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

package gtk

// #include <gtk/gtk.h>
import "C"

// TODO:
// gtk_widget_list_action_prefixes().
// gtk_widget_get_action_group().
func (v *Widget) GetActionGroup(name string) glib.IActionGroup{
  return C.gtk_widget_get_action_group(v.native(), (*C.gchar)(C.CString(name)))  
}
