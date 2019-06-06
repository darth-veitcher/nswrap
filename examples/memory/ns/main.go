package ns


/*
#cgo CFLAGS: -x objective-c -fno-objc-arc
#cgo LDFLAGS: -framework Foundation
#pragma clang diagnostic ignored "-Wformat-security"

#import <Foundation/Foundation.h>


void*
NSArray_AllocWithZone(void* zone) {
	return [NSArray allocWithZone:zone];
}
void* _Nonnull
NSArray_Array() {
	return [NSArray array];
}
BOOL
NSArray_AutomaticallyNotifiesObserversForKey(void* key) {
	return [NSArray automaticallyNotifiesObserversForKey:key];
}
void* _Nonnull
NSArray_KeyPathsForValuesAffectingValueForKey(void* key) {
	return [NSArray keyPathsForValuesAffectingValueForKey:key];
}
void*
selectorFromString(char *s) {
	return NSSelectorFromString([NSString stringWithUTF8String:s]);
}

BOOL
NSArray_InstancesRespondToSelector(void* aSelector) {
	return [NSArray instancesRespondToSelector:aSelector];
}
void*
NSArray_MutableCopyWithZone(void* zone) {
	return [NSArray mutableCopyWithZone:zone];
}
void* _Nonnull
NSArray_ClassForKeyedUnarchiver() {
	return [NSArray classForKeyedUnarchiver];
}
void*
NSArray_Description() {
	return [NSArray description];
}
void* _Nonnull
NSArray_ArrayWithArray(void* array) {
	return [NSArray arrayWithArray:array];
}
void*
NSArray_New() {
	return [NSArray new];
}
void*
NSArray_DebugDescription() {
	return [NSArray debugDescription];
}
NSInteger
NSArray_Version() {
	return [NSArray version];
}
BOOL
NSArray_ConformsToProtocol(void* protocol) {
	return [NSArray conformsToProtocol:protocol];
}
void*
NSArray_Superclass() {
	return [NSArray superclass];
}
void* _Nonnull
NSArray_ArrayWithObjectsCount(void* objects, NSUInteger cnt) {
	return [NSArray arrayWithObjects:objects count:cnt];
}
void* _Nonnull
NSArray_ArrayWithObjects(void* firstObj, void* object) {
	NSObject** arr = object;
	return [NSArray arrayWithObjects:firstObj , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
}
NSUInteger
NSArray_Hash() {
	return [NSArray hash];
}
void
NSArray_SetVersion(NSInteger aVersion) {
	[NSArray setVersion:aVersion];
}
BOOL
NSArray_ResolveInstanceMethod(void* sel) {
	return [NSArray resolveInstanceMethod:sel];
}
void*
NSArray_InstanceMethodSignatureForSelector(void* aSelector) {
	return [NSArray instanceMethodSignatureForSelector:aSelector];
}
void* _Nullable
NSArray_ArrayWithContentsOfURL(void* url) {
	return [NSArray arrayWithContentsOfURL:url];
}
void* _Nullable
NSArray_ArrayWithContentsOfURLError(void* url, void* error) {
	return [NSArray arrayWithContentsOfURL:url error:error];
}
BOOL
NSArray_ResolveClassMethod(void* sel) {
	return [NSArray resolveClassMethod:sel];
}
void
NSArray_CancelPreviousPerformRequestsWithTarget(void* aTarget) {
	[NSArray cancelPreviousPerformRequestsWithTarget:aTarget];
}
void
NSArray_CancelPreviousPerformRequestsWithTargetSelector(void* aTarget, void* aSelector, void* anArgument) {
	[NSArray cancelPreviousPerformRequestsWithTarget:aTarget selector:aSelector object:anArgument];
}
void
NSArray_Load() {
	[NSArray load];
}
void* _Nonnull
NSArray_ClassFallbacksForKeyedArchiver() {
	return [NSArray classFallbacksForKeyedArchiver];
}
void*
NSArray_CopyWithZone(void* zone) {
	return [NSArray copyWithZone:zone];
}
void*
NSArray_Alloc() {
	return [NSArray alloc];
}
BOOL
NSArray_AccessInstanceVariablesDirectly() {
	return [NSArray accessInstanceVariablesDirectly];
}
void*
NSArray_Class() {
	return [NSArray class];
}
void* _Nonnull
NSArray_ArrayWithObject(void* anObject) {
	return [NSArray arrayWithObject:anObject];
}
void* _Nullable
NSArray_ArrayWithContentsOfFile(void* path) {
	return [NSArray arrayWithContentsOfFile:path];
}
BOOL
NSArray_IsSubclassOfClass(void* aClass) {
	return [NSArray isSubclassOfClass:aClass];
}
void* _Nonnull
NSArray_inst_Description(void* o) {
	return [(NSArray*)o description];
}
void* _Nonnull
NSArray_inst_FilteredArrayUsingPredicate(void* o, void* predicate) {
	return [(NSArray*)o filteredArrayUsingPredicate:predicate];
}
void* _Nonnull
NSArray_inst_InitWithObjectsCount(void* o, void* objects, NSUInteger cnt) {
	return [(NSArray*)o initWithObjects:objects count:cnt];
}
void* _Nonnull
NSArray_inst_InitWithObjects(void* o, void* firstObj, void* object) {
	NSObject** arr = object;
	return [(NSArray*)o initWithObjects:firstObj , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
}
void* _Nonnull
NSArray_inst_ComponentsJoinedByString(void* o, void* separator) {
	return [(NSArray*)o componentsJoinedByString:separator];
}
void* _Nonnull
NSArray_inst_Init(void* o) {
	return [(NSArray*)o init];
}
void
NSArray_inst_SetValue(void* o, void* value, void* key) {
	[(NSArray*)o setValue:value forKey:key];
}
void* _Nonnull
NSArray_inst_ObjectEnumerator(void* o) {
	return [(NSArray*)o objectEnumerator];
}
void* _Nonnull
NSArray_inst_ObjectsAtIndexes(void* o, void* indexes) {
	return [(NSArray*)o objectsAtIndexes:indexes];
}
void* _Nonnull
NSArray_inst_SortedArrayHint(void* o) {
	return [(NSArray*)o sortedArrayHint];
}
void* _Nonnull
NSArray_inst_DescriptionWithLocale(void* o, void* locale) {
	return [(NSArray*)o descriptionWithLocale:locale];
}
void* _Nonnull
NSArray_inst_DescriptionWithLocaleIndent(void* o, void* locale, NSUInteger level) {
	return [(NSArray*)o descriptionWithLocale:locale indent:level];
}
NSUInteger
NSArray_inst_Count(void* o) {
	return [(NSArray*)o count];
}
void* _Nonnull
NSArray_inst_SortedArrayUsingSelector(void* o, void* comparator) {
	return [(NSArray*)o sortedArrayUsingSelector:comparator];
}
void* _Nullable
NSArray_inst_FirstObject(void* o) {
	return [(NSArray*)o firstObject];
}
void
NSArray_inst_MakeObjectsPerformSelector(void* o, void* aSelector) {
	[(NSArray*)o makeObjectsPerformSelector:aSelector];
}
void
NSArray_inst_MakeObjectsPerformSelectorWithObject(void* o, void* aSelector, void* argument) {
	[(NSArray*)o makeObjectsPerformSelector:aSelector withObject:argument];
}
void* _Nonnull
NSArray_inst_SubarrayWithRange(void* o, NSRange range) {
	return [(NSArray*)o subarrayWithRange:range];
}
void* _Nonnull
NSArray_inst_ArrayByAddingObject(void* o, void* anObject) {
	return [(NSArray*)o arrayByAddingObject:anObject];
}
void* _Nullable
NSArray_inst_InitWithContentsOfFile(void* o, void* path) {
	return [(NSArray*)o initWithContentsOfFile:path];
}
BOOL
NSArray_inst_WriteToURLError(void* o, void* url, void* error) {
	return [(NSArray*)o writeToURL:url error:error];
}
BOOL
NSArray_inst_WriteToURLAtomically(void* o, void* url, BOOL atomically) {
	return [(NSArray*)o writeToURL:url atomically:atomically];
}
void* _Nullable
NSArray_inst_InitWithContentsOfURL(void* o, void* url) {
	return [(NSArray*)o initWithContentsOfURL:url];
}
void* _Nullable
NSArray_inst_InitWithContentsOfURLError(void* o, void* url, void* error) {
	return [(NSArray*)o initWithContentsOfURL:url error:error];
}
void* _Nonnull
NSArray_inst_ArrayByAddingObjectsFromArray(void* o, void* otherArray) {
	return [(NSArray*)o arrayByAddingObjectsFromArray:otherArray];
}
void* _Nullable
NSArray_inst_FirstObjectCommonWithArray(void* o, void* otherArray) {
	return [(NSArray*)o firstObjectCommonWithArray:otherArray];
}
void* _Nonnull
NSArray_inst_InitWithArray(void* o, void* array) {
	return [(NSArray*)o initWithArray:array];
}
void* _Nonnull
NSArray_inst_InitWithArrayCopyItems(void* o, void* array, BOOL flag) {
	return [(NSArray*)o initWithArray:array copyItems:flag];
}
BOOL
NSArray_inst_ContainsObject(void* o, void* anObject) {
	return [(NSArray*)o containsObject:anObject];
}
void* _Nonnull
NSArray_inst_ObjectAtIndexedSubscript(void* o, NSUInteger idx) {
	return [(NSArray*)o objectAtIndexedSubscript:idx];
}
NSUInteger
NSArray_inst_IndexOfObjectIdenticalTo(void* o, void* anObject) {
	return [(NSArray*)o indexOfObjectIdenticalTo:anObject];
}
NSUInteger
NSArray_inst_IndexOfObjectIdenticalToInRange(void* o, void* anObject, NSRange range) {
	return [(NSArray*)o indexOfObjectIdenticalTo:anObject inRange:range];
}
BOOL
NSArray_inst_IsEqualToArray(void* o, void* otherArray) {
	return [(NSArray*)o isEqualToArray:otherArray];
}
void
NSArray_inst_AddObserverForKeyPath(void* o, void* observer, void* keyPath, NSKeyValueObservingOptions options, void* context) {
	[(NSArray*)o addObserver:observer forKeyPath:keyPath options:options context:context];
}
void
NSArray_inst_AddObserverToObjectsAtIndexes(void* o, void* observer, void* indexes, void* keyPath, NSKeyValueObservingOptions options, void* context) {
	[(NSArray*)o addObserver:observer toObjectsAtIndexes:indexes forKeyPath:keyPath options:options context:context];
}
NSUInteger
NSArray_inst_IndexOfObject(void* o, void* anObject) {
	return [(NSArray*)o indexOfObject:anObject];
}
NSUInteger
NSArray_inst_IndexOfObjectInRange(void* o, void* anObject, NSRange range) {
	return [(NSArray*)o indexOfObject:anObject inRange:range];
}
BOOL
NSArray_inst_WriteToFile(void* o, void* path, BOOL useAuxiliaryFile) {
	return [(NSArray*)o writeToFile:path atomically:useAuxiliaryFile];
}
void* _Nullable
NSArray_inst_InitWithCoder(void* o, void* aDecoder) {
	return [(NSArray*)o initWithCoder:aDecoder];
}
void* _Nonnull
NSArray_inst_ReverseObjectEnumerator(void* o) {
	return [(NSArray*)o reverseObjectEnumerator];
}
void* _Nonnull
NSArray_inst_ValueForKey(void* o, void* key) {
	return [(NSArray*)o valueForKey:key];
}
void* _Nonnull
NSArray_inst_PathsMatchingExtensions(void* o, void* filterTypes) {
	return [(NSArray*)o pathsMatchingExtensions:filterTypes];
}
void
NSArray_inst_GetObjects(void* o, void* objects, NSRange range) {
	[(NSArray*)o getObjects:objects range:range];
}
void* _Nonnull
NSArray_inst_ObjectAtIndex(void* o, NSUInteger index) {
	return [(NSArray*)o objectAtIndex:index];
}
void* _Nonnull
NSArray_inst_SortedArrayUsingDescriptors(void* o, void* sortDescriptors) {
	return [(NSArray*)o sortedArrayUsingDescriptors:sortDescriptors];
}
void* _Nullable
NSArray_inst_LastObject(void* o) {
	return [(NSArray*)o lastObject];
}
void
NSArray_inst_RemoveObserverForKeyPath(void* o, void* observer, void* keyPath) {
	[(NSArray*)o removeObserver:observer forKeyPath:keyPath];
}
void
NSArray_inst_RemoveObserverFromObjectsAtIndexes(void* o, void* observer, void* indexes, void* keyPath) {
	[(NSArray*)o removeObserver:observer fromObjectsAtIndexes:indexes forKeyPath:keyPath];
}
void
NSArray_inst_RemoveObserverForKeyPathContext(void* o, void* observer, void* keyPath, void* context) {
	[(NSArray*)o removeObserver:observer forKeyPath:keyPath context:context];
}
void
NSArray_inst_RemoveObserverFromObjectsAtIndexesForKeyPath(void* o, void* observer, void* indexes, void* keyPath, void* context) {
	[(NSArray*)o removeObserver:observer fromObjectsAtIndexes:indexes forKeyPath:keyPath context:context];
}
void* _Nonnull
NSArray_inst_CopyWithZone(void* o, void* zone) {
	return [(NSArray*)o copyWithZone:zone];
}
void* _Nonnull
NSArray_inst_MutableCopyWithZone(void* o, void* zone) {
	return [(NSArray*)o mutableCopyWithZone:zone];
}
BOOL
NSArray_SupportsSecureCoding() {
	return [NSArray supportsSecureCoding];
}
NSUInteger
NSArray_inst_CountByEnumeratingWithState(void* o, void* state, void* buffer, NSUInteger len) {
	return [(NSArray*)o countByEnumeratingWithState:state objects:buffer count:len];
}
void* _Nonnull
NSMutableArray_ArrayWithArray(void* array) {
	return [NSMutableArray arrayWithArray:array];
}
void* _Nonnull
NSMutableArray_KeyPathsForValuesAffectingValueForKey(void* key) {
	return [NSMutableArray keyPathsForValuesAffectingValueForKey:key];
}
void* _Nonnull
NSMutableArray_ClassFallbacksForKeyedArchiver() {
	return [NSMutableArray classFallbacksForKeyedArchiver];
}
void*
NSMutableArray_Description() {
	return [NSMutableArray description];
}
void* _Nonnull
NSMutableArray_Array() {
	return [NSMutableArray array];
}
void* _Nonnull
NSMutableArray_ArrayWithCapacity(NSUInteger numItems) {
	return [NSMutableArray arrayWithCapacity:numItems];
}
BOOL
NSMutableArray_AutomaticallyNotifiesObserversForKey(void* key) {
	return [NSMutableArray automaticallyNotifiesObserversForKey:key];
}
void
NSMutableArray_CancelPreviousPerformRequestsWithTarget(void* aTarget) {
	[NSMutableArray cancelPreviousPerformRequestsWithTarget:aTarget];
}
void
NSMutableArray_CancelPreviousPerformRequestsWithTargetSelector(void* aTarget, void* aSelector, void* anArgument) {
	[NSMutableArray cancelPreviousPerformRequestsWithTarget:aTarget selector:aSelector object:anArgument];
}
void*
NSMutableArray_MutableCopyWithZone(void* zone) {
	return [NSMutableArray mutableCopyWithZone:zone];
}
void*
NSMutableArray_Superclass() {
	return [NSMutableArray superclass];
}
void*
NSMutableArray_InstanceMethodSignatureForSelector(void* aSelector) {
	return [NSMutableArray instanceMethodSignatureForSelector:aSelector];
}
BOOL
NSMutableArray_ResolveClassMethod(void* sel) {
	return [NSMutableArray resolveClassMethod:sel];
}
void* _Nonnull
NSMutableArray_ArrayWithObject(void* anObject) {
	return [NSMutableArray arrayWithObject:anObject];
}
void
NSMutableArray_Load() {
	[NSMutableArray load];
}
void* _Nullable
NSMutableArray_ArrayWithContentsOfURL(void* url) {
	return [NSMutableArray arrayWithContentsOfURL:url];
}
void* _Nullable
NSMutableArray_ArrayWithContentsOfURLError(void* url, void* error) {
	return [NSMutableArray arrayWithContentsOfURL:url error:error];
}
BOOL
NSMutableArray_IsSubclassOfClass(void* aClass) {
	return [NSMutableArray isSubclassOfClass:aClass];
}
NSInteger
NSMutableArray_Version() {
	return [NSMutableArray version];
}
BOOL
NSMutableArray_ResolveInstanceMethod(void* sel) {
	return [NSMutableArray resolveInstanceMethod:sel];
}
void* _Nonnull
NSMutableArray_ClassForKeyedUnarchiver() {
	return [NSMutableArray classForKeyedUnarchiver];
}
void*
NSMutableArray_New() {
	return [NSMutableArray new];
}
BOOL
NSMutableArray_ConformsToProtocol(void* protocol) {
	return [NSMutableArray conformsToProtocol:protocol];
}
NSUInteger
NSMutableArray_Hash() {
	return [NSMutableArray hash];
}
void* _Nullable
NSMutableArray_ArrayWithContentsOfFile(void* path) {
	return [NSMutableArray arrayWithContentsOfFile:path];
}
void*
NSMutableArray_Class() {
	return [NSMutableArray class];
}
void*
NSMutableArray_CopyWithZone(void* zone) {
	return [NSMutableArray copyWithZone:zone];
}
void
NSMutableArray_SetVersion(NSInteger aVersion) {
	[NSMutableArray setVersion:aVersion];
}
void*
NSMutableArray_DebugDescription() {
	return [NSMutableArray debugDescription];
}
void*
NSMutableArray_AllocWithZone(void* zone) {
	return [NSMutableArray allocWithZone:zone];
}
BOOL
NSMutableArray_AccessInstanceVariablesDirectly() {
	return [NSMutableArray accessInstanceVariablesDirectly];
}
void*
NSMutableArray_Alloc() {
	return [NSMutableArray alloc];
}
BOOL
NSMutableArray_InstancesRespondToSelector(void* aSelector) {
	return [NSMutableArray instancesRespondToSelector:aSelector];
}
void* _Nonnull
NSMutableArray_ArrayWithObjectsCount(void* objects, NSUInteger cnt) {
	return [NSMutableArray arrayWithObjects:objects count:cnt];
}
void* _Nonnull
NSMutableArray_ArrayWithObjects(void* firstObj, void* object) {
	NSObject** arr = object;
	return [NSMutableArray arrayWithObjects:firstObj , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
}
void
NSMutableArray_inst_RemoveObject(void* o, void* anObject) {
	[(NSMutableArray*)o removeObject:anObject];
}
void
NSMutableArray_inst_RemoveObjectInRange(void* o, void* anObject, NSRange range) {
	[(NSMutableArray*)o removeObject:anObject inRange:range];
}
void
NSMutableArray_inst_RemoveLastObject(void* o) {
	[(NSMutableArray*)o removeLastObject];
}
void
NSMutableArray_inst_RemoveObjectsInArray(void* o, void* otherArray) {
	[(NSMutableArray*)o removeObjectsInArray:otherArray];
}
void* _Nullable
NSMutableArray_inst_InitWithCoder(void* o, void* aDecoder) {
	return [(NSMutableArray*)o initWithCoder:aDecoder];
}
void
NSMutableArray_inst_AddObject(void* o, void* anObject) {
	[(NSMutableArray*)o addObject:anObject];
}
void
NSMutableArray_inst_RemoveObjectIdenticalTo(void* o, void* anObject) {
	[(NSMutableArray*)o removeObjectIdenticalTo:anObject];
}
void
NSMutableArray_inst_RemoveObjectIdenticalToInRange(void* o, void* anObject, NSRange range) {
	[(NSMutableArray*)o removeObjectIdenticalTo:anObject inRange:range];
}
void
NSMutableArray_inst_ExchangeObjectAtIndex(void* o, NSUInteger idx1, NSUInteger idx2) {
	[(NSMutableArray*)o exchangeObjectAtIndex:idx1 withObjectAtIndex:idx2];
}
void
NSMutableArray_inst_RemoveObjectsAtIndexes(void* o, void* indexes) {
	[(NSMutableArray*)o removeObjectsAtIndexes:indexes];
}
void* _Nonnull
NSMutableArray_inst_Init(void* o) {
	return [(NSMutableArray*)o init];
}
void
NSMutableArray_inst_ReplaceObjectsInRangeWithObjectsFromArray(void* o, NSRange range, void* otherArray) {
	[(NSMutableArray*)o replaceObjectsInRange:range withObjectsFromArray:otherArray];
}
void
NSMutableArray_inst_ReplaceObjectsInRangeWithObjectsFromArrayRange(void* o, NSRange range, void* otherArray, NSRange otherRange) {
	[(NSMutableArray*)o replaceObjectsInRange:range withObjectsFromArray:otherArray range:otherRange];
}
void
NSMutableArray_inst_SetObject(void* o, void* obj, NSUInteger idx) {
	[(NSMutableArray*)o setObject:obj atIndexedSubscript:idx];
}
void
NSMutableArray_inst_InsertObject(void* o, void* anObject, NSUInteger index) {
	[(NSMutableArray*)o insertObject:anObject atIndex:index];
}
void
NSMutableArray_inst_InsertObjects(void* o, void* objects, void* indexes) {
	[(NSMutableArray*)o insertObjects:objects atIndexes:indexes];
}
void
NSMutableArray_inst_SortUsingDescriptors(void* o, void* sortDescriptors) {
	[(NSMutableArray*)o sortUsingDescriptors:sortDescriptors];
}
void
NSMutableArray_inst_RemoveAllObjects(void* o) {
	[(NSMutableArray*)o removeAllObjects];
}
void
NSMutableArray_inst_SetArray(void* o, void* otherArray) {
	[(NSMutableArray*)o setArray:otherArray];
}
void
NSMutableArray_inst_RemoveObjectsInRange(void* o, NSRange range) {
	[(NSMutableArray*)o removeObjectsInRange:range];
}
void
NSMutableArray_inst_ReplaceObjectAtIndex(void* o, NSUInteger index, void* anObject) {
	[(NSMutableArray*)o replaceObjectAtIndex:index withObject:anObject];
}
void
NSMutableArray_inst_SortUsingSelector(void* o, void* comparator) {
	[(NSMutableArray*)o sortUsingSelector:comparator];
}
void* _Nullable
NSMutableArray_inst_InitWithContentsOfFile(void* o, void* path) {
	return [(NSMutableArray*)o initWithContentsOfFile:path];
}
void
NSMutableArray_inst_FilterUsingPredicate(void* o, void* predicate) {
	[(NSMutableArray*)o filterUsingPredicate:predicate];
}
void* _Nonnull
NSMutableArray_inst_InitWithCapacity(void* o, NSUInteger numItems) {
	return [(NSMutableArray*)o initWithCapacity:numItems];
}
void
NSMutableArray_inst_ReplaceObjectsAtIndexes(void* o, void* indexes, void* objects) {
	[(NSMutableArray*)o replaceObjectsAtIndexes:indexes withObjects:objects];
}
void* _Nullable
NSMutableArray_inst_InitWithContentsOfURL(void* o, void* url) {
	return [(NSMutableArray*)o initWithContentsOfURL:url];
}
void
NSMutableArray_inst_AddObjectsFromArray(void* o, void* otherArray) {
	[(NSMutableArray*)o addObjectsFromArray:otherArray];
}
void
NSMutableArray_inst_RemoveObjectAtIndex(void* o, NSUInteger index) {
	[(NSMutableArray*)o removeObjectAtIndex:index];
}
void* _Nonnull
NSString_PathWithComponents(void* components) {
	return [NSString pathWithComponents:components];
}
void*
NSString_Description() {
	return [NSString description];
}
BOOL
NSString_AccessInstanceVariablesDirectly() {
	return [NSString accessInstanceVariablesDirectly];
}
NSUInteger
NSString_Hash() {
	return [NSString hash];
}
void* _Nonnull
NSString_StringWithString(void* string) {
	return [NSString stringWithString:string];
}
void* _Nullable
NSString_StringWithContentsOfURLEncoding(void* url, NSStringEncoding enc, void* error) {
	return [NSString stringWithContentsOfURL:url encoding:enc error:error];
}
void* _Nullable
NSString_StringWithContentsOfURLUsedEncoding(void* url, void* enc, void* error) {
	return [NSString stringWithContentsOfURL:url usedEncoding:enc error:error];
}
void* _Nullable
NSString_StringWithCString(void* cString, NSStringEncoding enc) {
	return [NSString stringWithCString:cString encoding:enc];
}
BOOL
NSString_InstancesRespondToSelector(void* aSelector) {
	return [NSString instancesRespondToSelector:aSelector];
}
void* _Nonnull
NSString_ClassForKeyedUnarchiver() {
	return [NSString classForKeyedUnarchiver];
}
void*
NSString_Class() {
	return [NSString class];
}
BOOL
NSString_ResolveClassMethod(void* sel) {
	return [NSString resolveClassMethod:sel];
}
void* _Nullable
NSString_StringWithContentsOfFileEncoding(void* path, NSStringEncoding enc, void* error) {
	return [NSString stringWithContentsOfFile:path encoding:enc error:error];
}
void* _Nullable
NSString_StringWithContentsOfFileUsedEncoding(void* path, void* enc, void* error) {
	return [NSString stringWithContentsOfFile:path usedEncoding:enc error:error];
}
void*
NSString_DebugDescription() {
	return [NSString debugDescription];
}
BOOL
NSString_ConformsToProtocol(void* protocol) {
	return [NSString conformsToProtocol:protocol];
}
NSStringEncoding
NSString_DefaultCStringEncoding() {
	return [NSString defaultCStringEncoding];
}
void*
NSString_MutableCopyWithZone(void* zone) {
	return [NSString mutableCopyWithZone:zone];
}
void*
NSString_CopyWithZone(void* zone) {
	return [NSString copyWithZone:zone];
}
void* _Nonnull
NSString_LocalizedNameOfStringEncoding(NSStringEncoding encoding) {
	return [NSString localizedNameOfStringEncoding:encoding];
}
void
NSString_CancelPreviousPerformRequestsWithTarget(void* aTarget) {
	[NSString cancelPreviousPerformRequestsWithTarget:aTarget];
}
void
NSString_CancelPreviousPerformRequestsWithTargetSelector(void* aTarget, void* aSelector, void* anArgument) {
	[NSString cancelPreviousPerformRequestsWithTarget:aTarget selector:aSelector object:anArgument];
}
NSInteger
NSString_Version() {
	return [NSString version];
}
BOOL
NSString_AutomaticallyNotifiesObserversForKey(void* key) {
	return [NSString automaticallyNotifiesObserversForKey:key];
}
const void* _Nonnull
NSString_AvailableStringEncodings() {
	return [NSString availableStringEncodings];
}
BOOL
NSString_IsSubclassOfClass(void* aClass) {
	return [NSString isSubclassOfClass:aClass];
}
void*
NSString_Superclass() {
	return [NSString superclass];
}
void*
NSString_New() {
	return [NSString new];
}
void
NSString_SetVersion(NSInteger aVersion) {
	[NSString setVersion:aVersion];
}
void*
NSString_AllocWithZone(void* zone) {
	return [NSString allocWithZone:zone];
}
void* _Nonnull
NSString_ClassFallbacksForKeyedArchiver() {
	return [NSString classFallbacksForKeyedArchiver];
}
void
NSString_Load() {
	[NSString load];
}
void* _Nonnull
NSString_LocalizedStringWithFormat(void* format, void* object) {
	NSObject** arr = object;
	return [NSString localizedStringWithFormat:format , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
}
void*
NSString_Alloc() {
	return [NSString alloc];
}
void* _Nonnull
NSString_StringWithCharacters(void* characters, NSUInteger length) {
	return [NSString stringWithCharacters:characters length:length];
}
NSStringEncoding
NSString_StringEncodingForData(void* data, void* opts, void* string, void* usedLossyConversion) {
	return [NSString stringEncodingForData:data encodingOptions:opts convertedString:string usedLossyConversion:usedLossyConversion];
}
BOOL
NSString_ResolveInstanceMethod(void* sel) {
	return [NSString resolveInstanceMethod:sel];
}
void* _Nullable
NSString_StringWithUTF8String(void* nullTerminatedCString) {
	return [NSString stringWithUTF8String:nullTerminatedCString];
}
void* _Nonnull
NSString_KeyPathsForValuesAffectingValueForKey(void* key) {
	return [NSString keyPathsForValuesAffectingValueForKey:key];
}
void* _Nonnull
NSString_String() {
	return [NSString string];
}
void*
NSString_InstanceMethodSignatureForSelector(void* aSelector) {
	return [NSString instanceMethodSignatureForSelector:aSelector];
}
void* _Nonnull
NSString_StringWithFormat(void* format, void* object) {
	NSObject** arr = object;
	return [NSString stringWithFormat:format , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
}
void* _Nonnull
NSString_inst_InitWithCharacters(void* o, void* characters, NSUInteger length) {
	return [(NSString*)o initWithCharacters:characters length:length];
}
void* _Nonnull
NSString_inst_StringByStandardizingPath(void* o) {
	return [(NSString*)o stringByStandardizingPath];
}
BOOL
NSString_inst_HasSuffix(void* o, void* str) {
	return [(NSString*)o hasSuffix:str];
}
void* _Nonnull
NSString_inst_ComponentsSeparatedByCharactersInSet(void* o, void* separator) {
	return [(NSString*)o componentsSeparatedByCharactersInSet:separator];
}
NSUInteger
NSString_inst_LengthOfBytesUsingEncoding(void* o, NSStringEncoding enc) {
	return [(NSString*)o lengthOfBytesUsingEncoding:enc];
}
BOOL
NSString_inst_LocalizedCaseInsensitiveContainsString(void* o, void* str) {
	return [(NSString*)o localizedCaseInsensitiveContainsString:str];
}
void* _Nonnull
NSString_inst_CapitalizedString(void* o) {
	return [(NSString*)o capitalizedString];
}
NSUInteger
NSString_inst_Hash(void* o) {
	return [(NSString*)o hash];
}
void* _Nullable
NSString_inst_InitWithContentsOfURLEncoding(void* o, void* url, NSStringEncoding enc, void* error) {
	return [(NSString*)o initWithContentsOfURL:url encoding:enc error:error];
}
void* _Nullable
NSString_inst_InitWithContentsOfURLUsedEncoding(void* o, void* url, void* enc, void* error) {
	return [(NSString*)o initWithContentsOfURL:url usedEncoding:enc error:error];
}
NSUInteger
NSString_inst_CompletePathIntoString(void* o, void* outputName, BOOL flag, void* outputArray, void* filterTypes) {
	return [(NSString*)o completePathIntoString:outputName caseSensitive:flag matchesIntoArray:outputArray filterTypes:filterTypes];
}
void* _Nonnull
NSString_inst_DecomposedStringWithCanonicalMapping(void* o) {
	return [(NSString*)o decomposedStringWithCanonicalMapping];
}
void* _Nonnull
NSString_inst_InitWithFormat(void* o, void* format, void* object) {
	NSObject** arr = object;
	return [(NSString*)o initWithFormat:format , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
}
void* _Nonnull
NSString_inst_InitWithFormatLocale(void* o, void* format, void* locale, void* object) {
	NSObject** arr = object;
	return [(NSString*)o initWithFormat:format locale:locale , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
}
void* _Nonnull
NSString_inst_SubstringWithRange(void* o, NSRange range) {
	return [(NSString*)o substringWithRange:range];
}
void
NSString_inst_GetCharacters(void* o, void* buffer) {
	[(NSString*)o getCharacters:buffer];
}
void
NSString_inst_GetCharactersRange(void* o, void* buffer, NSRange range) {
	[(NSString*)o getCharacters:buffer range:range];
}
void* _Nonnull
NSString_inst_LowercaseStringWithLocale(void* o, void* locale) {
	return [(NSString*)o lowercaseStringWithLocale:locale];
}
void* _Nonnull
NSString_inst_StringByExpandingTildeInPath(void* o) {
	return [(NSString*)o stringByExpandingTildeInPath];
}
void* _Nonnull
NSString_inst_StringByAppendingFormat(void* o, void* format, void* object) {
	NSObject** arr = object;
	return [(NSString*)o stringByAppendingFormat:format , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
}
NSRange
NSString_inst_RangeOfComposedCharacterSequencesForRange(void* o, NSRange range) {
	return [(NSString*)o rangeOfComposedCharacterSequencesForRange:range];
}
void* _Nullable
NSString_inst_PropertyListFromStringsFileFormat(void* o) {
	return [(NSString*)o propertyListFromStringsFileFormat];
}
void* _Nonnull
NSString_inst_CommonPrefixWithString(void* o, void* str, NSStringCompareOptions mask) {
	return [(NSString*)o commonPrefixWithString:str options:mask];
}
const void* _Nonnull
NSString_inst_FileSystemRepresentation(void* o) {
	return [(NSString*)o fileSystemRepresentation];
}
void* _Nonnull
NSString_inst_StringByFoldingWithOptions(void* o, NSStringCompareOptions options, void* locale) {
	return [(NSString*)o stringByFoldingWithOptions:options locale:locale];
}
void* _Nonnull
NSString_inst_StringsByAppendingPaths(void* o, void* paths) {
	return [(NSString*)o stringsByAppendingPaths:paths];
}
void* _Nonnull
NSString_inst_InitWithCharactersNoCopy(void* o, void* characters, NSUInteger length, BOOL freeBuffer) {
	return [(NSString*)o initWithCharactersNoCopy:characters length:length freeWhenDone:freeBuffer];
}
NSComparisonResult
NSString_inst_LocalizedStandardCompare(void* o, void* string) {
	return [(NSString*)o localizedStandardCompare:string];
}
void* _Nonnull
NSString_inst_LocalizedCapitalizedString(void* o) {
	return [(NSString*)o localizedCapitalizedString];
}
void* _Nonnull
NSString_inst_UppercaseString(void* o) {
	return [(NSString*)o uppercaseString];
}
void* _Nonnull
NSString_inst_PropertyList(void* o) {
	return [(NSString*)o propertyList];
}
NSRange
NSString_inst_LocalizedStandardRangeOfString(void* o, void* str) {
	return [(NSString*)o localizedStandardRangeOfString:str];
}
BOOL
NSString_inst_WriteToFile(void* o, void* path, BOOL useAuxiliaryFile, NSStringEncoding enc, void* error) {
	return [(NSString*)o writeToFile:path atomically:useAuxiliaryFile encoding:enc error:error];
}
unichar
NSString_inst_CharacterAtIndex(void* o, NSUInteger index) {
	return [(NSString*)o characterAtIndex:index];
}
void* _Nonnull
NSString_inst_StringByDeletingPathExtension(void* o) {
	return [(NSString*)o stringByDeletingPathExtension];
}
void* _Nonnull
NSString_inst_StringByTrimmingCharactersInSet(void* o, void* set) {
	return [(NSString*)o stringByTrimmingCharactersInSet:set];
}
void* _Nonnull
NSString_inst_PrecomposedStringWithCompatibilityMapping(void* o) {
	return [(NSString*)o precomposedStringWithCompatibilityMapping];
}
void* _Nonnull
NSString_inst_ComponentsSeparatedByString(void* o, void* separator) {
	return [(NSString*)o componentsSeparatedByString:separator];
}
void* _Nonnull
NSString_inst_StringByDeletingLastPathComponent(void* o) {
	return [(NSString*)o stringByDeletingLastPathComponent];
}
void* _Nonnull
NSString_inst_PrecomposedStringWithCanonicalMapping(void* o) {
	return [(NSString*)o precomposedStringWithCanonicalMapping];
}
void
NSString_inst_GetParagraphStart(void* o, void* startPtr, void* parEndPtr, void* contentsEndPtr, NSRange range) {
	[(NSString*)o getParagraphStart:startPtr end:parEndPtr contentsEnd:contentsEndPtr forRange:range];
}
void* _Nullable
NSString_inst_InitWithCoder(void* o, void* aDecoder) {
	return [(NSString*)o initWithCoder:aDecoder];
}
NSRange
NSString_inst_LineRangeForRange(void* o, NSRange range) {
	return [(NSString*)o lineRangeForRange:range];
}
NSRange
NSString_inst_RangeOfComposedCharacterSequenceAtIndex(void* o, NSUInteger index) {
	return [(NSString*)o rangeOfComposedCharacterSequenceAtIndex:index];
}
void* _Nonnull
NSString_inst_StringByAppendingPathComponent(void* o, void* str) {
	return [(NSString*)o stringByAppendingPathComponent:str];
}
BOOL
NSString_inst_WriteToURL(void* o, void* url, BOOL useAuxiliaryFile, NSStringEncoding enc, void* error) {
	return [(NSString*)o writeToURL:url atomically:useAuxiliaryFile encoding:enc error:error];
}
void* _Nullable
NSString_inst_InitWithBytes(void* o, void* bytes, NSUInteger len, NSStringEncoding encoding) {
	return [(NSString*)o initWithBytes:bytes length:len encoding:encoding];
}
void* _Nullable
NSString_inst_DataUsingEncoding(void* o, NSStringEncoding encoding) {
	return [(NSString*)o dataUsingEncoding:encoding];
}
void* _Nullable
NSString_inst_DataUsingEncodingAllowLossyConversion(void* o, NSStringEncoding encoding, BOOL lossy) {
	return [(NSString*)o dataUsingEncoding:encoding allowLossyConversion:lossy];
}
void
NSString_inst_GetLineStart(void* o, void* startPtr, void* lineEndPtr, void* contentsEndPtr, NSRange range) {
	[(NSString*)o getLineStart:startPtr end:lineEndPtr contentsEnd:contentsEndPtr forRange:range];
}
NSRange
NSString_inst_RangeOfCharacterFromSet(void* o, void* searchSet) {
	return [(NSString*)o rangeOfCharacterFromSet:searchSet];
}
NSRange
NSString_inst_RangeOfCharacterFromSetOptions(void* o, void* searchSet, NSStringCompareOptions mask) {
	return [(NSString*)o rangeOfCharacterFromSet:searchSet options:mask];
}
NSRange
NSString_inst_RangeOfCharacterFromSetOptionsRange(void* o, void* searchSet, NSStringCompareOptions mask, NSRange rangeOfReceiverToSearch) {
	return [(NSString*)o rangeOfCharacterFromSet:searchSet options:mask range:rangeOfReceiverToSearch];
}
void* _Nonnull
NSString_inst_LastPathComponent(void* o) {
	return [(NSString*)o lastPathComponent];
}
void* _Nonnull
NSString_inst_StringByResolvingSymlinksInPath(void* o) {
	return [(NSString*)o stringByResolvingSymlinksInPath];
}
void* _Nonnull
NSString_inst_PathExtension(void* o) {
	return [(NSString*)o pathExtension];
}
NSUInteger
NSString_inst_Length(void* o) {
	return [(NSString*)o length];
}
void* _Nonnull
NSString_inst_StringByReplacingCharactersInRange(void* o, NSRange range, void* replacement) {
	return [(NSString*)o stringByReplacingCharactersInRange:range withString:replacement];
}
NSUInteger
NSString_inst_MaximumLengthOfBytesUsingEncoding(void* o, NSStringEncoding enc) {
	return [(NSString*)o maximumLengthOfBytesUsingEncoding:enc];
}
void* _Nonnull
NSString_inst_UppercaseStringWithLocale(void* o, void* locale) {
	return [(NSString*)o uppercaseStringWithLocale:locale];
}
BOOL
NSString_inst_HasPrefix(void* o, void* str) {
	return [(NSString*)o hasPrefix:str];
}
void* _Nullable
NSString_inst_InitWithUTF8String(void* o, void* nullTerminatedCString) {
	return [(NSString*)o initWithUTF8String:nullTerminatedCString];
}
void* _Nonnull
NSString_inst_VariantFittingPresentationWidth(void* o, NSInteger width) {
	return [(NSString*)o variantFittingPresentationWidth:width];
}
void* _Nonnull
NSString_inst_SubstringFromIndex(void* o, NSUInteger from) {
	return [(NSString*)o substringFromIndex:from];
}
void* _Nonnull
NSString_inst_SubstringToIndex(void* o, NSUInteger to) {
	return [(NSString*)o substringToIndex:to];
}
NSComparisonResult
NSString_inst_Compare(void* o, void* string) {
	return [(NSString*)o compare:string];
}
NSComparisonResult
NSString_inst_CompareOptions(void* o, void* string, NSStringCompareOptions mask) {
	return [(NSString*)o compare:string options:mask];
}
NSComparisonResult
NSString_inst_CompareOptionsRange(void* o, void* string, NSStringCompareOptions mask, NSRange rangeOfReceiverToCompare) {
	return [(NSString*)o compare:string options:mask range:rangeOfReceiverToCompare];
}
NSComparisonResult
NSString_inst_CompareOptionsRangeLocale(void* o, void* string, NSStringCompareOptions mask, NSRange rangeOfReceiverToCompare, void* locale) {
	return [(NSString*)o compare:string options:mask range:rangeOfReceiverToCompare locale:locale];
}
NSInteger
NSString_inst_IntegerValue(void* o) {
	return [(NSString*)o integerValue];
}
void* _Nullable
NSString_inst_StringByRemovingPercentEncoding(void* o) {
	return [(NSString*)o stringByRemovingPercentEncoding];
}
void* _Nullable
NSString_inst_InitWithCString(void* o, void* nullTerminatedCString, NSStringEncoding encoding) {
	return [(NSString*)o initWithCString:nullTerminatedCString encoding:encoding];
}
BOOL
NSString_inst_IsEqualToString(void* o, void* aString) {
	return [(NSString*)o isEqualToString:aString];
}
void* _Nonnull
NSString_inst_LocalizedLowercaseString(void* o) {
	return [(NSString*)o localizedLowercaseString];
}
void* _Nonnull
NSString_inst_StringByAbbreviatingWithTildeInPath(void* o) {
	return [(NSString*)o stringByAbbreviatingWithTildeInPath];
}
void* _Nullable
NSString_inst_StringByApplyingTransform(void* o, void* transform, BOOL reverse) {
	return [(NSString*)o stringByApplyingTransform:transform reverse:reverse];
}
void* _Nonnull
NSString_inst_DecomposedStringWithCompatibilityMapping(void* o) {
	return [(NSString*)o decomposedStringWithCompatibilityMapping];
}
NSComparisonResult
NSString_inst_LocalizedCaseInsensitiveCompare(void* o, void* string) {
	return [(NSString*)o localizedCaseInsensitiveCompare:string];
}
NSStringEncoding
NSString_inst_FastestEncoding(void* o) {
	return [(NSString*)o fastestEncoding];
}
void* _Nonnull
NSString_inst_CapitalizedStringWithLocale(void* o, void* locale) {
	return [(NSString*)o capitalizedStringWithLocale:locale];
}
BOOL
NSString_inst_ContainsString(void* o, void* str) {
	return [(NSString*)o containsString:str];
}
void* _Nonnull
NSString_inst_InitWithString(void* o, void* aString) {
	return [(NSString*)o initWithString:aString];
}
BOOL
NSString_inst_BoolValue(void* o) {
	return [(NSString*)o boolValue];
}
NSStringEncoding
NSString_inst_SmallestEncoding(void* o) {
	return [(NSString*)o smallestEncoding];
}
NSComparisonResult
NSString_inst_LocalizedCompare(void* o, void* string) {
	return [(NSString*)o localizedCompare:string];
}
const void* _Nullable
NSString_inst_CStringUsingEncoding(void* o, NSStringEncoding encoding) {
	return [(NSString*)o cStringUsingEncoding:encoding];
}
BOOL
NSString_inst_LocalizedStandardContainsString(void* o, void* str) {
	return [(NSString*)o localizedStandardContainsString:str];
}
void* _Nullable
NSString_inst_StringByAppendingPathExtension(void* o, void* str) {
	return [(NSString*)o stringByAppendingPathExtension:str];
}
void* _Nullable
NSString_inst_InitWithBytesNoCopy(void* o, void* bytes, NSUInteger len, NSStringEncoding encoding, BOOL freeBuffer) {
	return [(NSString*)o initWithBytesNoCopy:bytes length:len encoding:encoding freeWhenDone:freeBuffer];
}
void* _Nonnull
NSString_inst_PathComponents(void* o) {
	return [(NSString*)o pathComponents];
}
double
NSString_inst_DoubleValue(void* o) {
	return [(NSString*)o doubleValue];
}
void* _Nullable
NSString_inst_InitWithContentsOfFileEncoding(void* o, void* path, NSStringEncoding enc, void* error) {
	return [(NSString*)o initWithContentsOfFile:path encoding:enc error:error];
}
void* _Nullable
NSString_inst_InitWithContentsOfFileUsedEncoding(void* o, void* path, void* enc, void* error) {
	return [(NSString*)o initWithContentsOfFile:path usedEncoding:enc error:error];
}
BOOL
NSString_inst_GetCString(void* o, void* buffer, NSUInteger maxBufferCount, NSStringEncoding encoding) {
	return [(NSString*)o getCString:buffer maxLength:maxBufferCount encoding:encoding];
}
void* _Nonnull
NSString_inst_StringByAppendingString(void* o, void* aString) {
	return [(NSString*)o stringByAppendingString:aString];
}
BOOL
NSString_inst_GetBytes(void* o, void* buffer, NSUInteger maxBufferCount, void* usedBufferCount, NSStringEncoding encoding, NSStringEncodingConversionOptions options, NSRange range, void* leftover) {
	return [(NSString*)o getBytes:buffer maxLength:maxBufferCount usedLength:usedBufferCount encoding:encoding options:options range:range remainingRange:leftover];
}
void* _Nonnull
NSString_inst_Init(void* o) {
	return [(NSString*)o init];
}
void* _Nullable
NSString_inst_InitWithData(void* o, void* data, NSStringEncoding encoding) {
	return [(NSString*)o initWithData:data encoding:encoding];
}
NSComparisonResult
NSString_inst_CaseInsensitiveCompare(void* o, void* string) {
	return [(NSString*)o caseInsensitiveCompare:string];
}
BOOL
NSString_inst_GetFileSystemRepresentation(void* o, void* cname, NSUInteger max) {
	return [(NSString*)o getFileSystemRepresentation:cname maxLength:max];
}
void* _Nonnull
NSString_inst_LowercaseString(void* o) {
	return [(NSString*)o lowercaseString];
}
void* _Nonnull
NSString_inst_StringByPaddingToLength(void* o, NSUInteger newLength, void* padString, NSUInteger padIndex) {
	return [(NSString*)o stringByPaddingToLength:newLength withString:padString startingAtIndex:padIndex];
}
void* _Nonnull
NSString_inst_StringByReplacingOccurrencesOfStringWithString(void* o, void* target, void* replacement) {
	return [(NSString*)o stringByReplacingOccurrencesOfString:target withString:replacement];
}
void* _Nonnull
NSString_inst_StringByReplacingOccurrencesOfStringWithStringOptions(void* o, void* target, void* replacement, NSStringCompareOptions options, NSRange searchRange) {
	return [(NSString*)o stringByReplacingOccurrencesOfString:target withString:replacement options:options range:searchRange];
}
int
NSString_inst_IntValue(void* o) {
	return [(NSString*)o intValue];
}
BOOL
NSString_inst_CanBeConvertedToEncoding(void* o, NSStringEncoding encoding) {
	return [(NSString*)o canBeConvertedToEncoding:encoding];
}
void* _Nullable
NSString_inst_StringByAddingPercentEncodingWithAllowedCharacters(void* o, void* allowedCharacters) {
	return [(NSString*)o stringByAddingPercentEncodingWithAllowedCharacters:allowedCharacters];
}
const void* _Nullable
NSString_inst_UTF8String(void* o) {
	return [(NSString*)o UTF8String];
}
NSRange
NSString_inst_ParagraphRangeForRange(void* o, NSRange range) {
	return [(NSString*)o paragraphRangeForRange:range];
}
void* _Nonnull
NSString_inst_LinguisticTagsInRange(void* o, NSRange range, void* scheme, NSLinguisticTaggerOptions options, void* orthography, void* tokenRanges) {
	return [(NSString*)o linguisticTagsInRange:range scheme:scheme options:options orthography:orthography tokenRanges:tokenRanges];
}
void* _Nonnull
NSString_inst_Description(void* o) {
	return [(NSString*)o description];
}
float
NSString_inst_FloatValue(void* o) {
	return [(NSString*)o floatValue];
}
void* _Nonnull
NSString_inst_LocalizedUppercaseString(void* o) {
	return [(NSString*)o localizedUppercaseString];
}
long long
NSString_inst_LongLongValue(void* o) {
	return [(NSString*)o longLongValue];
}
BOOL
NSString_inst_IsAbsolutePath(void* o) {
	return [(NSString*)o isAbsolutePath];
}
NSRange
NSString_inst_RangeOfString(void* o, void* searchString) {
	return [(NSString*)o rangeOfString:searchString];
}
NSRange
NSString_inst_RangeOfStringOptions(void* o, void* searchString, NSStringCompareOptions mask) {
	return [(NSString*)o rangeOfString:searchString options:mask];
}
NSRange
NSString_inst_RangeOfStringOptionsRange(void* o, void* searchString, NSStringCompareOptions mask, NSRange rangeOfReceiverToSearch) {
	return [(NSString*)o rangeOfString:searchString options:mask range:rangeOfReceiverToSearch];
}
NSRange
NSString_inst_RangeOfStringOptionsRangeLocale(void* o, void* searchString, NSStringCompareOptions mask, NSRange rangeOfReceiverToSearch, void* locale) {
	return [(NSString*)o rangeOfString:searchString options:mask range:rangeOfReceiverToSearch locale:locale];
}
void* _Nonnull
NSString_inst_CopyWithZone(void* o, void* zone) {
	return [(NSString*)o copyWithZone:zone];
}
void* _Nonnull
NSString_inst_MutableCopyWithZone(void* o, void* zone) {
	return [(NSString*)o mutableCopyWithZone:zone];
}
BOOL
NSString_SupportsSecureCoding() {
	return [NSString supportsSecureCoding];
}
void* _Nullable
NSString_ObjectWithItemProviderData(void* data, void* typeIdentifier, void* outError) {
	return [NSString objectWithItemProviderData:data typeIdentifier:typeIdentifier error:outError];
}
void* _Nonnull
NSString_ReadableTypeIdentifiersForItemProvider() {
	return [NSString readableTypeIdentifiersForItemProvider];
}
NSItemProviderRepresentationVisibility
NSString_ItemProviderVisibilityForRepresentationWithTypeIdentifier(void* typeIdentifier) {
	return [NSString itemProviderVisibilityForRepresentationWithTypeIdentifier:typeIdentifier];
}
void* _Nonnull
NSString_WritableTypeIdentifiersForItemProvider() {
	return [NSString writableTypeIdentifiersForItemProvider];
}
NSItemProviderRepresentationVisibility
NSString_inst_ItemProviderVisibilityForRepresentationWithTypeIdentifier(void* o, void* typeIdentifier) {
	return [(NSString*)o itemProviderVisibilityForRepresentationWithTypeIdentifier:typeIdentifier];
}
void* _Nonnull
NSString_inst_WritableTypeIdentifiersForItemProvider(void* o) {
	return [(NSString*)o writableTypeIdentifiersForItemProvider];
}
void*
NSObject_Description() {
	return [NSObject description];
}
void*
NSObject_CopyWithZone(void* zone) {
	return [NSObject copyWithZone:zone];
}
void*
NSObject_MutableCopyWithZone(void* zone) {
	return [NSObject mutableCopyWithZone:zone];
}
void
NSObject_CancelPreviousPerformRequestsWithTarget(void* aTarget) {
	[NSObject cancelPreviousPerformRequestsWithTarget:aTarget];
}
void
NSObject_CancelPreviousPerformRequestsWithTargetSelector(void* aTarget, void* aSelector, void* anArgument) {
	[NSObject cancelPreviousPerformRequestsWithTarget:aTarget selector:aSelector object:anArgument];
}
NSUInteger
NSObject_Hash() {
	return [NSObject hash];
}
void*
NSObject_AllocWithZone(void* zone) {
	return [NSObject allocWithZone:zone];
}
void*
NSObject_DebugDescription() {
	return [NSObject debugDescription];
}
BOOL
NSObject_IsSubclassOfClass(void* aClass) {
	return [NSObject isSubclassOfClass:aClass];
}
BOOL
NSObject_InstancesRespondToSelector(void* aSelector) {
	return [NSObject instancesRespondToSelector:aSelector];
}
void* _Nonnull
NSObject_KeyPathsForValuesAffectingValueForKey(void* key) {
	return [NSObject keyPathsForValuesAffectingValueForKey:key];
}
BOOL
NSObject_ResolveClassMethod(void* sel) {
	return [NSObject resolveClassMethod:sel];
}
void*
NSObject_Alloc() {
	return [NSObject alloc];
}
void* _Nonnull
NSObject_ClassFallbacksForKeyedArchiver() {
	return [NSObject classFallbacksForKeyedArchiver];
}
void* _Nonnull
NSObject_ClassForKeyedUnarchiver() {
	return [NSObject classForKeyedUnarchiver];
}
void
NSObject_SetVersion(NSInteger aVersion) {
	[NSObject setVersion:aVersion];
}
NSInteger
NSObject_Version() {
	return [NSObject version];
}
BOOL
NSObject_ConformsToProtocol(void* protocol) {
	return [NSObject conformsToProtocol:protocol];
}
void*
NSObject_InstanceMethodSignatureForSelector(void* aSelector) {
	return [NSObject instanceMethodSignatureForSelector:aSelector];
}
void*
NSObject_Superclass() {
	return [NSObject superclass];
}
BOOL
NSObject_AutomaticallyNotifiesObserversForKey(void* key) {
	return [NSObject automaticallyNotifiesObserversForKey:key];
}
void
NSObject_Load() {
	[NSObject load];
}
BOOL
NSObject_AccessInstanceVariablesDirectly() {
	return [NSObject accessInstanceVariablesDirectly];
}
void*
NSObject_New() {
	return [NSObject new];
}
BOOL
NSObject_ResolveInstanceMethod(void* sel) {
	return [NSObject resolveInstanceMethod:sel];
}
void*
NSObject_Class() {
	return [NSObject class];
}
void* _Nonnull
NSObject_inst_AutoContentAccessingProxy(void* o) {
	return [(NSObject*)o autoContentAccessingProxy];
}
BOOL
NSObject_inst_IsLessThanOrEqualTo(void* o, void* object) {
	return [(NSObject*)o isLessThanOrEqualTo:object];
}
void*
NSObject_inst_MutableCopy(void* o) {
	return [(NSObject*)o mutableCopy];
}
void
NSObject_inst_DoesNotRecognizeSelector(void* o, void* aSelector) {
	[(NSObject*)o doesNotRecognizeSelector:aSelector];
}
void* _Nonnull
NSObject_inst_DictionaryWithValuesForKeys(void* o, void* keys) {
	return [(NSObject*)o dictionaryWithValuesForKeys:keys];
}
void* _Nullable
NSObject_inst_ObservationInfo(void* o) {
	return [(NSObject*)o observationInfo];
}
BOOL
NSObject_inst_AttemptRecoveryFromErrorOptionIndex(void* o, void* error, NSUInteger recoveryOptionIndex) {
	return [(NSObject*)o attemptRecoveryFromError:error optionIndex:recoveryOptionIndex];
}
void
NSObject_inst_AttemptRecoveryFromErrorOptionIndexDelegate(void* o, void* error, NSUInteger recoveryOptionIndex, void* delegate, void* didRecoverSelector, void* contextInfo) {
	[(NSObject*)o attemptRecoveryFromError:error optionIndex:recoveryOptionIndex delegate:delegate didRecoverSelector:didRecoverSelector contextInfo:contextInfo];
}
void
NSObject_inst_InsertValueInPropertyWithKey(void* o, void* value, void* key) {
	[(NSObject*)o insertValue:value inPropertyWithKey:key];
}
void
NSObject_inst_InsertValueAtIndex(void* o, void* value, NSUInteger index, void* key) {
	[(NSObject*)o insertValue:value atIndex:index inPropertyWithKey:key];
}
void
NSObject_inst_WillChange(void* o, NSKeyValueChange changeKind, void* indexes, void* key) {
	[(NSObject*)o willChange:changeKind valuesAtIndexes:indexes forKey:key];
}
void* _Nonnull
NSObject_inst_MutableSetValueForKeyPath(void* o, void* keyPath) {
	return [(NSObject*)o mutableSetValueForKeyPath:keyPath];
}
void
NSObject_inst_ObserveValueForKeyPath(void* o, void* keyPath, void* object, void* change, void* context) {
	[(NSObject*)o observeValueForKeyPath:keyPath ofObject:object change:change context:context];
}
BOOL
NSObject_inst_ScriptingBeginsWith(void* o, void* object) {
	return [(NSObject*)o scriptingBeginsWith:object];
}
void* _Nonnull
NSObject_inst_MutableArrayValueForKeyPath(void* o, void* keyPath) {
	return [(NSObject*)o mutableArrayValueForKeyPath:keyPath];
}
void* _Nonnull
NSObject_inst_AttributeKeys(void* o) {
	return [(NSObject*)o attributeKeys];
}
void
NSObject_inst_SetValueForKey(void* o, void* value, void* key) {
	[(NSObject*)o setValue:value forKey:key];
}
void
NSObject_inst_SetValueForKeyPath(void* o, void* value, void* keyPath) {
	[(NSObject*)o setValue:value forKeyPath:keyPath];
}
void
NSObject_inst_SetValueForUndefinedKey(void* o, void* value, void* key) {
	[(NSObject*)o setValue:value forUndefinedKey:key];
}
void* _Nullable
NSObject_inst_ValueForKeyPath(void* o, void* keyPath) {
	return [(NSObject*)o valueForKeyPath:keyPath];
}
void* _Nullable
NSObject_inst_CoerceValue(void* o, void* value, void* key) {
	return [(NSObject*)o coerceValue:value forKey:key];
}
void
NSObject_inst_DidChangeValueForKey(void* o, void* key) {
	[(NSObject*)o didChangeValueForKey:key];
}
void
NSObject_inst_DidChangeValueForKeyWithSetMutation(void* o, void* key, NSKeyValueSetMutationKind mutationKind, void* objects) {
	[(NSObject*)o didChangeValueForKey:key withSetMutation:mutationKind usingObjects:objects];
}
BOOL
NSObject_inst_ValidateValueForKey(void* o, void* ioValue, void* inKey, void* outError) {
	return [(NSObject*)o validateValue:ioValue forKey:inKey error:outError];
}
BOOL
NSObject_inst_ValidateValueForKeyPath(void* o, void* ioValue, void* inKeyPath, void* outError) {
	return [(NSObject*)o validateValue:ioValue forKeyPath:inKeyPath error:outError];
}
void* _Nullable
NSObject_inst_InverseForRelationshipKey(void* o, void* relationshipKey) {
	return [(NSObject*)o inverseForRelationshipKey:relationshipKey];
}
void* _Nullable
NSObject_inst_CopyScriptingValue(void* o, void* value, void* key, void* properties) {
	return [(NSObject*)o copyScriptingValue:value forKey:key withProperties:properties];
}
void* _Nullable
NSObject_inst_ValueForKey(void* o, void* key) {
	return [(NSObject*)o valueForKey:key];
}
void
NSObject_inst_AddObserver(void* o, void* observer, void* keyPath, NSKeyValueObservingOptions options, void* context) {
	[(NSObject*)o addObserver:observer forKeyPath:keyPath options:options context:context];
}
void* _Nonnull
NSObject_inst_ClassDescription(void* o) {
	return [(NSObject*)o classDescription];
}
void* _Nullable
NSObject_inst_ValueAtIndex(void* o, NSUInteger index, void* key) {
	return [(NSObject*)o valueAtIndex:index inPropertyWithKey:key];
}
void* _Nonnull
NSObject_inst_MutableOrderedSetValueForKeyPath(void* o, void* keyPath) {
	return [(NSObject*)o mutableOrderedSetValueForKeyPath:keyPath];
}
void
NSObject_inst_DidChange(void* o, NSKeyValueChange changeKind, void* indexes, void* key) {
	[(NSObject*)o didChange:changeKind valuesAtIndexes:indexes forKey:key];
}
BOOL
NSObject_inst_IsGreaterThanOrEqualTo(void* o, void* object) {
	return [(NSObject*)o isGreaterThanOrEqualTo:object];
}
void* _Nonnull
NSObject_inst_MutableOrderedSetValueForKey(void* o, void* key) {
	return [(NSObject*)o mutableOrderedSetValueForKey:key];
}
void
NSObject_inst_PerformSelectorWithObject(void* o, void* aSelector, void* anArgument, NSTimeInterval delay) {
	[(NSObject*)o performSelector:aSelector withObject:anArgument afterDelay:delay];
}
void
NSObject_inst_PerformSelectorWithObjectAfterDelay(void* o, void* aSelector, void* anArgument, NSTimeInterval delay, void* modes) {
	[(NSObject*)o performSelector:aSelector withObject:anArgument afterDelay:delay inModes:modes];
}
void
NSObject_inst_PerformSelectorOnThread(void* o, void* aSelector, void* thr, void* arg, BOOL wait) {
	[(NSObject*)o performSelector:aSelector onThread:thr withObject:arg waitUntilDone:wait];
}
void
NSObject_inst_PerformSelectorOnThreadWithObject(void* o, void* aSelector, void* thr, void* arg, BOOL wait, void* array) {
	[(NSObject*)o performSelector:aSelector onThread:thr withObject:arg waitUntilDone:wait modes:array];
}
BOOL
NSObject_inst_IsCaseInsensitiveLike(void* o, void* object) {
	return [(NSObject*)o isCaseInsensitiveLike:object];
}
FourCharCode
NSObject_inst_ClassCode(void* o) {
	return [(NSObject*)o classCode];
}
void* _Nonnull
NSObject_inst_MutableSetValueForKey(void* o, void* key) {
	return [(NSObject*)o mutableSetValueForKey:key];
}
void
NSObject_inst_SetValuesForKeysWithDictionary(void* o, void* keyedValues) {
	[(NSObject*)o setValuesForKeysWithDictionary:keyedValues];
}
void
NSObject_inst_RemoveObserverForKeyPath(void* o, void* observer, void* keyPath) {
	[(NSObject*)o removeObserver:observer forKeyPath:keyPath];
}
void
NSObject_inst_RemoveObserverForKeyPathContext(void* o, void* observer, void* keyPath, void* context) {
	[(NSObject*)o removeObserver:observer forKeyPath:keyPath context:context];
}
void* _Nonnull
NSObject_inst_ClassName(void* o) {
	return [(NSObject*)o className];
}
void
NSObject_inst_ForwardInvocation(void* o, void* anInvocation) {
	[(NSObject*)o forwardInvocation:anInvocation];
}
void* _Nullable
NSObject_inst_ClassForKeyedArchiver(void* o) {
	return [(NSObject*)o classForKeyedArchiver];
}
void* _Nonnull
NSObject_inst_ToManyRelationshipKeys(void* o) {
	return [(NSObject*)o toManyRelationshipKeys];
}
void* _Nullable
NSObject_inst_ObjectSpecifier(void* o) {
	return [(NSObject*)o objectSpecifier];
}
void*
NSObject_inst_MethodSignatureForSelector(void* o, void* aSelector) {
	return [(NSObject*)o methodSignatureForSelector:aSelector];
}
BOOL
NSObject_inst_ScriptingIsLessThanOrEqualTo(void* o, void* object) {
	return [(NSObject*)o scriptingIsLessThanOrEqualTo:object];
}
void*
NSObject_inst_Copy(void* o) {
	return [(NSObject*)o copy];
}
void* _Nullable
NSObject_inst_ReplacementObjectForCoder(void* o, void* aCoder) {
	return [(NSObject*)o replacementObjectForCoder:aCoder];
}
void
NSObject_inst_PerformSelectorInBackground(void* o, void* aSelector, void* arg) {
	[(NSObject*)o performSelectorInBackground:aSelector withObject:arg];
}
BOOL
NSObject_inst_ScriptingIsGreaterThanOrEqualTo(void* o, void* object) {
	return [(NSObject*)o scriptingIsGreaterThanOrEqualTo:object];
}
BOOL
NSObject_inst_ScriptingIsGreaterThan(void* o, void* object) {
	return [(NSObject*)o scriptingIsGreaterThan:object];
}
void
NSObject_inst_Dealloc(void* o) {
	[(NSObject*)o dealloc];
}
BOOL
NSObject_inst_ScriptingEndsWith(void* o, void* object) {
	return [(NSObject*)o scriptingEndsWith:object];
}
void* _Nonnull
NSObject_inst_ToOneRelationshipKeys(void* o) {
	return [(NSObject*)o toOneRelationshipKeys];
}
BOOL
NSObject_inst_IsLike(void* o, void* object) {
	return [(NSObject*)o isLike:object];
}
void
NSObject_inst_PerformSelectorOnMainThreadWithObject(void* o, void* aSelector, void* arg, BOOL wait) {
	[(NSObject*)o performSelectorOnMainThread:aSelector withObject:arg waitUntilDone:wait];
}
void
NSObject_inst_PerformSelectorOnMainThreadWithObjectWaitUntilDone(void* o, void* aSelector, void* arg, BOOL wait, void* array) {
	[(NSObject*)o performSelectorOnMainThread:aSelector withObject:arg waitUntilDone:wait modes:array];
}
void*
NSObject_inst_ForwardingTargetForSelector(void* o, void* aSelector) {
	return [(NSObject*)o forwardingTargetForSelector:aSelector];
}
void
NSObject_inst_SetNilValueForKey(void* o, void* key) {
	[(NSObject*)o setNilValueForKey:key];
}
void* _Nullable
NSObject_inst_IndicesOfObjectsByEvaluatingObjectSpecifier(void* o, void* specifier) {
	return [(NSObject*)o indicesOfObjectsByEvaluatingObjectSpecifier:specifier];
}
void* _Nullable
NSObject_inst_ReplacementObjectForKeyedArchiver(void* o, void* archiver) {
	return [(NSObject*)o replacementObjectForKeyedArchiver:archiver];
}
BOOL
NSObject_inst_ScriptingContains(void* o, void* object) {
	return [(NSObject*)o scriptingContains:object];
}
BOOL
NSObject_inst_IsLessThan(void* o, void* object) {
	return [(NSObject*)o isLessThan:object];
}
void*
NSObject_inst_Init(void* o) {
	return [(NSObject*)o init];
}
void
NSObject_inst_RemoveValueAtIndex(void* o, NSUInteger index, void* key) {
	[(NSObject*)o removeValueAtIndex:index fromPropertyWithKey:key];
}
void
NSObject_inst_WillChangeValueForKey(void* o, void* key) {
	[(NSObject*)o willChangeValueForKey:key];
}
void
NSObject_inst_WillChangeValueForKeyWithSetMutation(void* o, void* key, NSKeyValueSetMutationKind mutationKind, void* objects) {
	[(NSObject*)o willChangeValueForKey:key withSetMutation:mutationKind usingObjects:objects];
}
void* _Nullable
NSObject_inst_ValueForUndefinedKey(void* o, void* key) {
	return [(NSObject*)o valueForUndefinedKey:key];
}
void
NSObject_inst_ReplaceValueAtIndex(void* o, NSUInteger index, void* key, void* value) {
	[(NSObject*)o replaceValueAtIndex:index inPropertyWithKey:key withValue:value];
}
void* _Nullable
NSObject_inst_ScriptingProperties(void* o) {
	return [(NSObject*)o scriptingProperties];
}
void* _Nullable
NSObject_inst_ValueWithName(void* o, void* name, void* key) {
	return [(NSObject*)o valueWithName:name inPropertyWithKey:key];
}
void* _Nullable
NSObject_inst_AwakeAfterUsingCoder(void* o, void* aDecoder) {
	return [(NSObject*)o awakeAfterUsingCoder:aDecoder];
}
void
NSObject_inst_SetScriptingProperties(void* o, void* scriptingProperties) {
	[(NSObject*)o setScriptingProperties:scriptingProperties];
}
void* _Nonnull
NSObject_inst_ClassForCoder(void* o) {
	return [(NSObject*)o classForCoder];
}
void* _Nullable
NSObject_inst_ValueWithUniqueID(void* o, void* uniqueID, void* key) {
	return [(NSObject*)o valueWithUniqueID:uniqueID inPropertyWithKey:key];
}
BOOL
NSObject_inst_ScriptingIsLessThan(void* o, void* object) {
	return [(NSObject*)o scriptingIsLessThan:object];
}
BOOL
NSObject_inst_IsEqualTo(void* o, void* object) {
	return [(NSObject*)o isEqualTo:object];
}
void
NSObject_inst_SetObservationInfo(void* o, void* observationInfo) {
	[(NSObject*)o setObservationInfo:observationInfo];
}
BOOL
NSObject_inst_IsGreaterThan(void* o, void* object) {
	return [(NSObject*)o isGreaterThan:object];
}
void* _Nonnull
NSObject_inst_MutableArrayValueForKey(void* o, void* key) {
	return [(NSObject*)o mutableArrayValueForKey:key];
}
BOOL
NSObject_inst_DoesContain(void* o, void* object) {
	return [(NSObject*)o doesContain:object];
}
void* _Nullable
NSObject_inst_NewScriptingObjectOfClass(void* o, void* objectClass, void* key, void* contentsValue, void* properties) {
	return [(NSObject*)o newScriptingObjectOfClass:objectClass forValueForKey:key withContentsValue:contentsValue properties:properties];
}
BOOL
NSObject_inst_ScriptingIsEqualTo(void* o, void* object) {
	return [(NSObject*)o scriptingIsEqualTo:object];
}
BOOL
NSObject_inst_IsNotEqualTo(void* o, void* object) {
	return [(NSObject*)o isNotEqualTo:object];
}
void* _Nullable
NSObject_inst_ScriptingValueForSpecifier(void* o, void* objectSpecifier) {
	return [(NSObject*)o scriptingValueForSpecifier:objectSpecifier];
}
void* _Nullable
NSObject_inst_ClassForArchiver(void* o) {
	return [(NSObject*)o classForArchiver];
}
BOOL
NSObject_inst_ConformsToProtocol(void* o, void* aProtocol) {
	return [(NSObject*)o conformsToProtocol:aProtocol];
}
void*
NSObject_inst_Description(void* o) {
	return [(NSObject*)o description];
}
void*
NSObject_inst_DebugDescription(void* o) {
	return [(NSObject*)o debugDescription];
}
NSUInteger
NSObject_inst_Hash(void* o) {
	return [(NSObject*)o hash];
}
void*
NSObject_inst_Class(void* o) {
	return [(NSObject*)o class];
}
void*
NSObject_inst_Self(void* o) {
	return [(NSObject*)o self];
}
BOOL
NSObject_inst_IsProxy(void* o) {
	return [(NSObject*)o isProxy];
}
BOOL
NSObject_inst_IsMemberOfClass(void* o, void* aClass) {
	return [(NSObject*)o isMemberOfClass:aClass];
}
BOOL
NSObject_inst_RespondsToSelector(void* o, void* aSelector) {
	return [(NSObject*)o respondsToSelector:aSelector];
}
BOOL
NSObject_inst_IsEqual(void* o, void* object) {
	return [(NSObject*)o isEqual:object];
}
void*
NSObject_inst_PerformSelector(void* o, void* aSelector) {
	return [(NSObject*)o performSelector:aSelector];
}
void*
NSObject_inst_PerformSelectorWithObjectWithObject(void* o, void* aSelector, void* object1, void* object2) {
	return [(NSObject*)o performSelector:aSelector withObject:object1 withObject:object2];
}
void*
NSObject_inst_Retain(void* o) {
	return [(NSObject*)o retain];
}
void
NSObject_inst_Release(void* o) {
	[(NSObject*)o release];
}
BOOL
NSObject_inst_IsKindOfClass(void* o, void* aClass) {
	return [(NSObject*)o isKindOfClass:aClass];
}
void*
NSObject_inst_Autorelease(void* o) {
	return [(NSObject*)o autorelease];
}
NSUInteger
NSObject_inst_RetainCount(void* o) {
	return [(NSObject*)o retainCount];
}
void*
NSObject_inst_Zone(void* o) {
	return [(NSObject*)o zone];
}
void*
NSObject_inst_Superclass(void* o) {
	return [(NSObject*)o superclass];
}
void* _Nonnull
NSAutoreleasePool_init(void* o) {
	return [(NSAutoreleasePool*)o init];
}

BOOL
NSAutoreleasePool_ConformsToProtocol(void* protocol) {
	return [NSAutoreleasePool conformsToProtocol:protocol];
}
NSInteger
NSAutoreleasePool_Version() {
	return [NSAutoreleasePool version];
}
void*
NSAutoreleasePool_New() {
	return [NSAutoreleasePool new];
}
BOOL
NSAutoreleasePool_ResolveClassMethod(void* sel) {
	return [NSAutoreleasePool resolveClassMethod:sel];
}
void*
NSAutoreleasePool_AllocWithZone(void* zone) {
	return [NSAutoreleasePool allocWithZone:zone];
}
void
NSAutoreleasePool_AddObject(void* anObject) {
	[NSAutoreleasePool addObject:anObject];
}
void*
NSAutoreleasePool_DebugDescription() {
	return [NSAutoreleasePool debugDescription];
}
void*
NSAutoreleasePool_Description() {
	return [NSAutoreleasePool description];
}
void
NSAutoreleasePool_SetVersion(NSInteger aVersion) {
	[NSAutoreleasePool setVersion:aVersion];
}
void*
NSAutoreleasePool_CopyWithZone(void* zone) {
	return [NSAutoreleasePool copyWithZone:zone];
}
void*
NSAutoreleasePool_MutableCopyWithZone(void* zone) {
	return [NSAutoreleasePool mutableCopyWithZone:zone];
}
BOOL
NSAutoreleasePool_IsSubclassOfClass(void* aClass) {
	return [NSAutoreleasePool isSubclassOfClass:aClass];
}
void*
NSAutoreleasePool_InstanceMethodSignatureForSelector(void* aSelector) {
	return [NSAutoreleasePool instanceMethodSignatureForSelector:aSelector];
}
NSUInteger
NSAutoreleasePool_Hash() {
	return [NSAutoreleasePool hash];
}
void*
NSAutoreleasePool_Superclass() {
	return [NSAutoreleasePool superclass];
}
BOOL
NSAutoreleasePool_ResolveInstanceMethod(void* sel) {
	return [NSAutoreleasePool resolveInstanceMethod:sel];
}
BOOL
NSAutoreleasePool_InstancesRespondToSelector(void* aSelector) {
	return [NSAutoreleasePool instancesRespondToSelector:aSelector];
}
void*
NSAutoreleasePool_Class() {
	return [NSAutoreleasePool class];
}
void
NSAutoreleasePool_Load() {
	[NSAutoreleasePool load];
}
void*
NSAutoreleasePool_Alloc() {
	return [NSAutoreleasePool alloc];
}
void
NSAutoreleasePool_inst_AddObject(void* o, void* anObject) {
	[(NSAutoreleasePool*)o addObject:anObject];
}
void
NSAutoreleasePool_inst_Drain(void* o) {
	[(NSAutoreleasePool*)o drain];
}

void MyClassDealloc(void*);
void MyClassRelease(void*);

@interface MyClass : NSObject 
{ }
- (void)dealloc;
- (void)release;
- (void)super_release;

@end
void MyClass_super_release(void* o);

@implementation MyClass

- (void)dealloc
{
	MyClassDealloc(self);
	[super dealloc];
}


- (void)release
{
	MyClassRelease(self);
}


- (void)super_release
{
	[super release];
}


@end

void MyClass_super_release(void* o)
{
	[(MyClass*)o super_release];
}


void*
MyClassAlloc() {
	return [MyClass alloc];
}

*/
import "C"

import (
	"unsafe"
)

type Id struct {
	ptr unsafe.Pointer
}
func (o Id) Ptr() unsafe.Pointer { return o.ptr }

type NSObject interface {
	Ptr() unsafe.Pointer
}

type NSString struct { Id }
func (o NSString) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSString() NSString {
	ret := NSString{}
	ret.ptr = o.ptr
	return ret
}

type NSArray struct { Id }
func (o NSArray) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSArray() NSArray {
	ret := NSArray{}
	ret.ptr = o.ptr
	return ret
}

type _NSZone C.struct__NSZone

type BOOL C.uchar

type NSSet struct { Id }
func (o NSSet) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSSet() NSSet {
	ret := NSSet{}
	ret.ptr = o.ptr
	return ret
}

type SEL *C.struct_objc_selector

type Class *C.struct_objc_class

type NSInteger C.long

type Protocol interface {
	Ptr() unsafe.Pointer
}

type NSUInteger C.ulong

type NSMethodSignature struct { Id }
func (o NSMethodSignature) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSMethodSignature() NSMethodSignature {
	ret := NSMethodSignature{}
	ret.ptr = o.ptr
	return ret
}

type NSURL struct { Id }
func (o NSURL) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSURL() NSURL {
	ret := NSURL{}
	ret.ptr = o.ptr
	return ret
}

type NSError struct { Id }
func (o NSError) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSError() NSError {
	ret := NSError{}
	ret.ptr = o.ptr
	return ret
}

type NSPredicate struct { Id }
func (o NSPredicate) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSPredicate() NSPredicate {
	ret := NSPredicate{}
	ret.ptr = o.ptr
	return ret
}

type NSEnumerator struct { Id }
func (o NSEnumerator) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSEnumerator() NSEnumerator {
	ret := NSEnumerator{}
	ret.ptr = o.ptr
	return ret
}

type NSIndexSet struct { Id }
func (o NSIndexSet) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSIndexSet() NSIndexSet {
	ret := NSIndexSet{}
	ret.ptr = o.ptr
	return ret
}

type NSData struct { Id }
func (o NSData) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSData() NSData {
	ret := NSData{}
	ret.ptr = o.ptr
	return ret
}

type NSRange C.struct__NSRange

type NSKeyValueObservingOptions C.enum_NSKeyValueObservingOptions

type NSCoder struct { Id }
func (o NSCoder) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSCoder() NSCoder {
	ret := NSCoder{}
	ret.ptr = o.ptr
	return ret
}

type NSZone C.struct__NSZone

type NSFastEnumerationState C.struct_NSFastEnumerationState

type NSMutableArray struct { NSArray }
func (o NSMutableArray) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSMutableArray() NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = o.ptr
	return ret
}

type NSStringEncoding C.NSUInteger

type Char C.char

type Unichar C.ushort

type NSDictionary struct { Id }
func (o NSDictionary) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSDictionary() NSDictionary {
	ret := NSDictionary{}
	ret.ptr = o.ptr
	return ret
}

type NSCharacterSet struct { Id }
func (o NSCharacterSet) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSCharacterSet() NSCharacterSet {
	ret := NSCharacterSet{}
	ret.ptr = o.ptr
	return ret
}

type NSLocale struct { Id }
func (o NSLocale) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSLocale() NSLocale {
	ret := NSLocale{}
	ret.ptr = o.ptr
	return ret
}

type NSStringCompareOptions C.enum_NSStringCompareOptions

type NSComparisonResult C.enum_NSComparisonResult

type NSStringTransform struct { NSString }
func (o NSStringTransform) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSStringTransform() NSStringTransform {
	ret := NSStringTransform{}
	ret.ptr = o.ptr
	return ret
}

type Double C.double

type NSStringEncodingConversionOptions C.enum_NSStringEncodingConversionOptions

type NSRangePointer *C.NSRange

type Int C.int

type NSLinguisticTagScheme struct { NSString }
func (o NSLinguisticTagScheme) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSLinguisticTagScheme() NSLinguisticTagScheme {
	ret := NSLinguisticTagScheme{}
	ret.ptr = o.ptr
	return ret
}

type NSLinguisticTaggerOptions C.enum_NSLinguisticTaggerOptions

type NSOrthography struct { Id }
func (o NSOrthography) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSOrthography() NSOrthography {
	ret := NSOrthography{}
	ret.ptr = o.ptr
	return ret
}

type Float C.float

type LongLong C.longlong

type NSItemProviderRepresentationVisibility C.enum_NSItemProviderRepresentationVisibility

type NSKeyValueChange C.enum_NSKeyValueChange

type NSMutableSet struct { NSSet }
func (o NSMutableSet) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSMutableSet() NSMutableSet {
	ret := NSMutableSet{}
	ret.ptr = o.ptr
	return ret
}

type NSKeyValueSetMutationKind C.enum_NSKeyValueSetMutationKind

type NSClassDescription struct { Id }
func (o NSClassDescription) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSClassDescription() NSClassDescription {
	ret := NSClassDescription{}
	ret.ptr = o.ptr
	return ret
}

type NSOrderedSet struct { Id }
func (o NSOrderedSet) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSOrderedSet() NSOrderedSet {
	ret := NSOrderedSet{}
	ret.ptr = o.ptr
	return ret
}

type NSMutableOrderedSet struct { NSOrderedSet }
func (o NSMutableOrderedSet) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSMutableOrderedSet() NSMutableOrderedSet {
	ret := NSMutableOrderedSet{}
	ret.ptr = o.ptr
	return ret
}

type NSTimeInterval C.double

type NSThread struct { Id }
func (o NSThread) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSThread() NSThread {
	ret := NSThread{}
	ret.ptr = o.ptr
	return ret
}

type FourCharCode C.UInt32

type NSInvocation struct { Id }
func (o NSInvocation) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSInvocation() NSInvocation {
	ret := NSInvocation{}
	ret.ptr = o.ptr
	return ret
}

type NSScriptObjectSpecifier struct { Id }
func (o NSScriptObjectSpecifier) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSScriptObjectSpecifier() NSScriptObjectSpecifier {
	ret := NSScriptObjectSpecifier{}
	ret.ptr = o.ptr
	return ret
}

type NSKeyedArchiver struct { NSCoder }
func (o NSKeyedArchiver) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSKeyedArchiver() NSKeyedArchiver {
	ret := NSKeyedArchiver{}
	ret.ptr = o.ptr
	return ret
}

type NSAutoreleasePool struct { Id }
func (o NSAutoreleasePool) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSAutoreleasePool() NSAutoreleasePool {
	ret := NSAutoreleasePool{}
	ret.ptr = o.ptr
	return ret
}

type MyClass struct { Id }
func (o MyClass) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) MyClass() MyClass {
	ret := MyClass{}
	ret.ptr = o.ptr
	return ret
}

func Selector(s string) SEL {
	return (SEL)(unsafe.Pointer(C.selectorFromString(C.CString(s))))
}

func (o NSString) String() string {
	return o.UTF8String().String()
}

func CharWithGoString(s string) *Char {
	return (*Char)(unsafe.Pointer(C.CString(s)))
}

func CharWithBytes(b []byte) *Char {
	return (*Char)(unsafe.Pointer(C.CString(string(b))))
}

func (c *Char) String() string {
	return C.GoString((*C.char)(c))
}

func (c *Char) Free() {
	C.free(unsafe.Pointer(c))
}

func (o NSAutoreleasePool) Init() NSAutoreleasePool {
	ret := NSAutoreleasePool{}
	ret.ptr = C.NSAutoreleasePool_init(o.Ptr())
	return ret
}

func Autoreleasepool(f func()) {
	pool := NSAutoreleasePoolAlloc().Init()
	f()
	pool.Drain()
}

func NSArrayAllocWithZone(zone *_NSZone) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_AllocWithZone(unsafe.Pointer(zone))
	return ret
}

func NSArrayArray() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_Array()
	return ret
}

func NSArrayAutomaticallyNotifiesObserversForKey(key NSString) bool {
	ret := (C.NSArray_AutomaticallyNotifiesObserversForKey(key.Ptr())) != 0
	return ret
}

func NSArrayKeyPathsForValuesAffectingValueForKey(key NSString) NSSet {
	ret := NSSet{}
	ret.ptr = C.NSArray_KeyPathsForValuesAffectingValueForKey(key.Ptr())
	return ret
}

func NSArrayInstancesRespondToSelector(aSelector SEL) bool {
	ret := (C.NSArray_InstancesRespondToSelector(unsafe.Pointer(aSelector))) != 0
	return ret
}

func NSArrayMutableCopyWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSArray_MutableCopyWithZone(unsafe.Pointer(zone))
	return ret
}

func NSArrayClassForKeyedUnarchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSArray_ClassForKeyedUnarchiver()))
	return ret
}

func NSArrayDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSArray_Description()
	return ret
}

func NSArrayWithArray(array NSArray) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_ArrayWithArray(array.Ptr())
	return ret
}

func NSArrayNew() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_New()
	return ret
}

func NSArrayDebugDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSArray_DebugDescription()
	return ret
}

func NSArrayVersion() NSInteger {
	ret := (NSInteger)(C.NSArray_Version())
	return ret
}

func NSArrayConformsToProtocol(protocol Protocol) bool {
	ret := (C.NSArray_ConformsToProtocol(protocol.Ptr())) != 0
	return ret
}

func NSArraySuperclass() Class {
	ret := (Class)(unsafe.Pointer(C.NSArray_Superclass()))
	return ret
}

func NSArrayWithObjectsCount(objects *[]Id, cnt NSUInteger) NSArray {

	goSlice0 := make([]unsafe.Pointer,cap(*objects))
	for i := 0; i < len(*objects); i++ {
		goSlice0[i] = (*objects)[i].Ptr()
	}
	ret := NSArray{}
	ret.ptr = C.NSArray_ArrayWithObjectsCount(unsafe.Pointer(&goSlice0[0]), (C.NSUInteger)(cnt))
	(*objects) = (*objects)[:cap(*objects)]
	for i := 0; i < len(*objects); i++ {
		if goSlice0[i] == nil {
			(*objects) = (*objects)[:i]
			break
		}
		(*objects)[i].ptr = goSlice0[i]
	}
	return ret
}

func NSArrayWithObjects(firstObj NSObject, objects ...NSObject) NSArray {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := NSArray{}
	ret.ptr = C.NSArray_ArrayWithObjects(firstObj.Ptr(), unsafe.Pointer(&object))
	return ret
}

func NSArrayHash() NSUInteger {
	ret := (NSUInteger)(C.NSArray_Hash())
	return ret
}

func NSArraySetVersion(aVersion NSInteger)  {
	C.NSArray_SetVersion((C.NSInteger)(aVersion))
}

func NSArrayResolveInstanceMethod(sel SEL) bool {
	ret := (C.NSArray_ResolveInstanceMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSArrayInstanceMethodSignatureForSelector(aSelector SEL) NSMethodSignature {
	ret := NSMethodSignature{}
	ret.ptr = C.NSArray_InstanceMethodSignatureForSelector(unsafe.Pointer(aSelector))
	return ret
}

func NSArrayWithContentsOfURL(url NSURL) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_ArrayWithContentsOfURL(url.Ptr())
	return ret
}

func NSArrayWithContentsOfURLError(url NSURL, error *[]NSError) NSArray {

	goSlice1 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice1[i] = (*error)[i].Ptr()
	}
	ret := NSArray{}
	ret.ptr = C.NSArray_ArrayWithContentsOfURLError(url.Ptr(), unsafe.Pointer(&goSlice1[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice1[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice1[i]
	}
	return ret
}

func NSArrayResolveClassMethod(sel SEL) bool {
	ret := (C.NSArray_ResolveClassMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSArrayCancelPreviousPerformRequestsWithTarget(aTarget NSObject)  {
	C.NSArray_CancelPreviousPerformRequestsWithTarget(aTarget.Ptr())
}

func NSArrayCancelPreviousPerformRequestsWithTargetSelector(aTarget NSObject, aSelector SEL, anArgument NSObject)  {
	C.NSArray_CancelPreviousPerformRequestsWithTargetSelector(aTarget.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr())
}

func NSArrayLoad()  {
	C.NSArray_Load()
}

func NSArrayClassFallbacksForKeyedArchiver() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_ClassFallbacksForKeyedArchiver()
	return ret
}

func NSArrayCopyWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSArray_CopyWithZone(unsafe.Pointer(zone))
	return ret
}

func NSArrayAlloc() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_Alloc()
	return ret
}

func NSArrayAccessInstanceVariablesDirectly() bool {
	ret := (C.NSArray_AccessInstanceVariablesDirectly()) != 0
	return ret
}

func NSArrayClass() Class {
	ret := (Class)(unsafe.Pointer(C.NSArray_Class()))
	return ret
}

func NSArrayWithObject(anObject NSObject) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_ArrayWithObject(anObject.Ptr())
	return ret
}

func NSArrayWithContentsOfFile(path NSString) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_ArrayWithContentsOfFile(path.Ptr())
	return ret
}

func NSArrayIsSubclassOfClass(aClass Class) bool {
	ret := (C.NSArray_IsSubclassOfClass(unsafe.Pointer(aClass))) != 0
	return ret
}

func (o NSArray) Description() NSString {
	ret := NSString{}
	ret.ptr = C.NSArray_inst_Description(o.Ptr())
	return ret
}

func (o NSArray) FilteredArrayUsingPredicate(predicate NSPredicate) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_FilteredArrayUsingPredicate(o.Ptr(), predicate.Ptr())
	return ret
}

func (o NSArray) InitWithObjectsCount(objects *[]Id, cnt NSUInteger) NSArray {

	goSlice1 := make([]unsafe.Pointer,cap(*objects))
	for i := 0; i < len(*objects); i++ {
		goSlice1[i] = (*objects)[i].Ptr()
	}
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_InitWithObjectsCount(o.Ptr(), unsafe.Pointer(&goSlice1[0]), (C.NSUInteger)(cnt))
	(*objects) = (*objects)[:cap(*objects)]
	for i := 0; i < len(*objects); i++ {
		if goSlice1[i] == nil {
			(*objects) = (*objects)[:i]
			break
		}
		(*objects)[i].ptr = goSlice1[i]
	}
	return ret
}

func (o NSArray) InitWithObjects(firstObj NSObject, objects ...NSObject) NSArray {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_InitWithObjects(o.Ptr(), firstObj.Ptr(), unsafe.Pointer(&object))
	return ret
}

func (o NSArray) ComponentsJoinedByString(separator NSString) NSString {
	ret := NSString{}
	ret.ptr = C.NSArray_inst_ComponentsJoinedByString(o.Ptr(), separator.Ptr())
	return ret
}

func (o NSArray) Init() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_Init(o.Ptr())
	return ret
}

func (o NSArray) SetValue(value NSObject, key NSString)  {
	C.NSArray_inst_SetValue(o.Ptr(), value.Ptr(), key.Ptr())
}

func (o NSArray) ObjectEnumerator() NSEnumerator {
	ret := NSEnumerator{}
	ret.ptr = C.NSArray_inst_ObjectEnumerator(o.Ptr())
	return ret
}

func (o NSArray) ObjectsAtIndexes(indexes NSIndexSet) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_ObjectsAtIndexes(o.Ptr(), indexes.Ptr())
	return ret
}

func (o NSArray) SortedArrayHint() NSData {
	ret := NSData{}
	ret.ptr = C.NSArray_inst_SortedArrayHint(o.Ptr())
	return ret
}

func (o NSArray) DescriptionWithLocale(locale NSObject) NSString {
	ret := NSString{}
	ret.ptr = C.NSArray_inst_DescriptionWithLocale(o.Ptr(), locale.Ptr())
	return ret
}

func (o NSArray) DescriptionWithLocaleIndent(locale NSObject, level NSUInteger) NSString {
	ret := NSString{}
	ret.ptr = C.NSArray_inst_DescriptionWithLocaleIndent(o.Ptr(), locale.Ptr(), (C.NSUInteger)(level))
	return ret
}

func (o NSArray) Count() NSUInteger {
	ret := (NSUInteger)(C.NSArray_inst_Count(o.Ptr()))
	return ret
}

func (o NSArray) SortedArrayUsingSelector(comparator SEL) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_SortedArrayUsingSelector(o.Ptr(), unsafe.Pointer(comparator))
	return ret
}

func (o NSArray) FirstObject() Id {
	ret := Id{}
	ret.ptr = C.NSArray_inst_FirstObject(o.Ptr())
	return ret
}

func (o NSArray) MakeObjectsPerformSelector(aSelector SEL)  {
	C.NSArray_inst_MakeObjectsPerformSelector(o.Ptr(), unsafe.Pointer(aSelector))
}

func (o NSArray) MakeObjectsPerformSelectorWithObject(aSelector SEL, argument NSObject)  {
	C.NSArray_inst_MakeObjectsPerformSelectorWithObject(o.Ptr(), unsafe.Pointer(aSelector), argument.Ptr())
}

func (o NSArray) SubarrayWithRange(range_ NSRange) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_SubarrayWithRange(o.Ptr(), (C.NSRange)(range_))
	return ret
}

func (o NSArray) ArrayByAddingObject(anObject NSObject) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_ArrayByAddingObject(o.Ptr(), anObject.Ptr())
	return ret
}

func (o NSArray) InitWithContentsOfFile(path NSString) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_InitWithContentsOfFile(o.Ptr(), path.Ptr())
	return ret
}

func (o NSArray) WriteToURLError(url NSURL, error *[]NSError) bool {

	goSlice2 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice2[i] = (*error)[i].Ptr()
	}
	ret := (C.NSArray_inst_WriteToURLError(o.Ptr(), url.Ptr(), unsafe.Pointer(&goSlice2[0]))) != 0
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice2[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice2[i]
	}
	return ret
}

func (o NSArray) WriteToURLAtomically(url NSURL, atomically BOOL) bool {
	ret := (C.NSArray_inst_WriteToURLAtomically(o.Ptr(), url.Ptr(), (C.BOOL)(atomically))) != 0
	return ret
}

func (o NSArray) InitWithContentsOfURL(url NSURL) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_InitWithContentsOfURL(o.Ptr(), url.Ptr())
	return ret
}

func (o NSArray) InitWithContentsOfURLError(url NSURL, error *[]NSError) NSArray {

	goSlice2 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice2[i] = (*error)[i].Ptr()
	}
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_InitWithContentsOfURLError(o.Ptr(), url.Ptr(), unsafe.Pointer(&goSlice2[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice2[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice2[i]
	}
	return ret
}

func (o NSArray) ArrayByAddingObjectsFromArray(otherArray NSArray) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_ArrayByAddingObjectsFromArray(o.Ptr(), otherArray.Ptr())
	return ret
}

func (o NSArray) FirstObjectCommonWithArray(otherArray NSArray) Id {
	ret := Id{}
	ret.ptr = C.NSArray_inst_FirstObjectCommonWithArray(o.Ptr(), otherArray.Ptr())
	return ret
}

func (o NSArray) InitWithArray(array NSArray) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_InitWithArray(o.Ptr(), array.Ptr())
	return ret
}

func (o NSArray) InitWithArrayCopyItems(array NSArray, flag BOOL) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_InitWithArrayCopyItems(o.Ptr(), array.Ptr(), (C.BOOL)(flag))
	return ret
}

func (o NSArray) ContainsObject(anObject NSObject) bool {
	ret := (C.NSArray_inst_ContainsObject(o.Ptr(), anObject.Ptr())) != 0
	return ret
}

func (o NSArray) ObjectAtIndexedSubscript(idx NSUInteger) Id {
	ret := Id{}
	ret.ptr = C.NSArray_inst_ObjectAtIndexedSubscript(o.Ptr(), (C.NSUInteger)(idx))
	return ret
}

func (o NSArray) IndexOfObjectIdenticalTo(anObject NSObject) NSUInteger {
	ret := (NSUInteger)(C.NSArray_inst_IndexOfObjectIdenticalTo(o.Ptr(), anObject.Ptr()))
	return ret
}

func (o NSArray) IndexOfObjectIdenticalToInRange(anObject NSObject, range_ NSRange) NSUInteger {
	ret := (NSUInteger)(C.NSArray_inst_IndexOfObjectIdenticalToInRange(o.Ptr(), anObject.Ptr(), (C.NSRange)(range_)))
	return ret
}

func (o NSArray) IsEqualToArray(otherArray NSArray) bool {
	ret := (C.NSArray_inst_IsEqualToArray(o.Ptr(), otherArray.Ptr())) != 0
	return ret
}

func (o NSArray) AddObserverForKeyPath(observer NSObject, keyPath NSString, options NSKeyValueObservingOptions, context unsafe.Pointer)  {
	C.NSArray_inst_AddObserverForKeyPath(o.Ptr(), observer.Ptr(), keyPath.Ptr(), (C.NSKeyValueObservingOptions)(options), unsafe.Pointer(context))
}

func (o NSArray) AddObserverToObjectsAtIndexes(observer NSObject, indexes NSIndexSet, keyPath NSString, options NSKeyValueObservingOptions, context unsafe.Pointer)  {
	C.NSArray_inst_AddObserverToObjectsAtIndexes(o.Ptr(), observer.Ptr(), indexes.Ptr(), keyPath.Ptr(), (C.NSKeyValueObservingOptions)(options), unsafe.Pointer(context))
}

func (o NSArray) IndexOfObject(anObject NSObject) NSUInteger {
	ret := (NSUInteger)(C.NSArray_inst_IndexOfObject(o.Ptr(), anObject.Ptr()))
	return ret
}

func (o NSArray) IndexOfObjectInRange(anObject NSObject, range_ NSRange) NSUInteger {
	ret := (NSUInteger)(C.NSArray_inst_IndexOfObjectInRange(o.Ptr(), anObject.Ptr(), (C.NSRange)(range_)))
	return ret
}

func (o NSArray) WriteToFile(path NSString, useAuxiliaryFile BOOL) bool {
	ret := (C.NSArray_inst_WriteToFile(o.Ptr(), path.Ptr(), (C.BOOL)(useAuxiliaryFile))) != 0
	return ret
}

func (o NSArray) InitWithCoder(aDecoder NSCoder) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_InitWithCoder(o.Ptr(), aDecoder.Ptr())
	return ret
}

func (o NSArray) ReverseObjectEnumerator() NSEnumerator {
	ret := NSEnumerator{}
	ret.ptr = C.NSArray_inst_ReverseObjectEnumerator(o.Ptr())
	return ret
}

func (o NSArray) ValueForKey(key NSString) Id {
	ret := Id{}
	ret.ptr = C.NSArray_inst_ValueForKey(o.Ptr(), key.Ptr())
	return ret
}

func (o NSArray) PathsMatchingExtensions(filterTypes NSArray) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_PathsMatchingExtensions(o.Ptr(), filterTypes.Ptr())
	return ret
}

func (o NSArray) GetObjects(objects *[]Id, range_ NSRange)  {

	goSlice1 := make([]unsafe.Pointer,cap(*objects))
	for i := 0; i < len(*objects); i++ {
		goSlice1[i] = (*objects)[i].Ptr()
	}
	C.NSArray_inst_GetObjects(o.Ptr(), unsafe.Pointer(&goSlice1[0]), (C.NSRange)(range_))
	(*objects) = (*objects)[:cap(*objects)]
	for i := 0; i < len(*objects); i++ {
		if goSlice1[i] == nil {
			(*objects) = (*objects)[:i]
			break
		}
		(*objects)[i].ptr = goSlice1[i]
	}
}

func (o NSArray) ObjectAtIndex(index NSUInteger) Id {
	ret := Id{}
	ret.ptr = C.NSArray_inst_ObjectAtIndex(o.Ptr(), (C.NSUInteger)(index))
	return ret
}

func (o NSArray) SortedArrayUsingDescriptors(sortDescriptors NSArray) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSArray_inst_SortedArrayUsingDescriptors(o.Ptr(), sortDescriptors.Ptr())
	return ret
}

func (o NSArray) LastObject() Id {
	ret := Id{}
	ret.ptr = C.NSArray_inst_LastObject(o.Ptr())
	return ret
}

func (o NSArray) RemoveObserverForKeyPath(observer NSObject, keyPath NSString)  {
	C.NSArray_inst_RemoveObserverForKeyPath(o.Ptr(), observer.Ptr(), keyPath.Ptr())
}

func (o NSArray) RemoveObserverFromObjectsAtIndexes(observer NSObject, indexes NSIndexSet, keyPath NSString)  {
	C.NSArray_inst_RemoveObserverFromObjectsAtIndexes(o.Ptr(), observer.Ptr(), indexes.Ptr(), keyPath.Ptr())
}

func (o NSArray) RemoveObserverForKeyPathContext(observer NSObject, keyPath NSString, context unsafe.Pointer)  {
	C.NSArray_inst_RemoveObserverForKeyPathContext(o.Ptr(), observer.Ptr(), keyPath.Ptr(), unsafe.Pointer(context))
}

func (o NSArray) RemoveObserverFromObjectsAtIndexesForKeyPath(observer NSObject, indexes NSIndexSet, keyPath NSString, context unsafe.Pointer)  {
	C.NSArray_inst_RemoveObserverFromObjectsAtIndexesForKeyPath(o.Ptr(), observer.Ptr(), indexes.Ptr(), keyPath.Ptr(), unsafe.Pointer(context))
}

func (o NSArray) CopyWithZone(zone *NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSArray_inst_CopyWithZone(o.Ptr(), unsafe.Pointer(zone))
	return ret
}

func (o NSArray) MutableCopyWithZone(zone *NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSArray_inst_MutableCopyWithZone(o.Ptr(), unsafe.Pointer(zone))
	return ret
}

func NSArraySupportsSecureCoding() bool {
	ret := (C.NSArray_SupportsSecureCoding()) != 0
	return ret
}

func (o NSArray) CountByEnumeratingWithState(state *NSFastEnumerationState, buffer *[]Id, len_ NSUInteger) NSUInteger {

	goSlice2 := make([]unsafe.Pointer,cap(*buffer))
	for i := 0; i < len(*buffer); i++ {
		goSlice2[i] = (*buffer)[i].Ptr()
	}
	ret := (NSUInteger)(C.NSArray_inst_CountByEnumeratingWithState(o.Ptr(), unsafe.Pointer(state), unsafe.Pointer(&goSlice2[0]), (C.NSUInteger)(len_)))
	(*buffer) = (*buffer)[:cap(*buffer)]
	for i := 0; i < len(*buffer); i++ {
		if goSlice2[i] == nil {
			(*buffer) = (*buffer)[:i]
			break
		}
		(*buffer)[i].ptr = goSlice2[i]
	}
	return ret
}

func NSMutableArrayWithArray(array NSArray) NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_ArrayWithArray(array.Ptr())
	return ret
}

func NSMutableArrayKeyPathsForValuesAffectingValueForKey(key NSString) NSSet {
	ret := NSSet{}
	ret.ptr = C.NSMutableArray_KeyPathsForValuesAffectingValueForKey(key.Ptr())
	return ret
}

func NSMutableArrayClassFallbacksForKeyedArchiver() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSMutableArray_ClassFallbacksForKeyedArchiver()
	return ret
}

func NSMutableArrayDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSMutableArray_Description()
	return ret
}

func NSMutableArrayArray() NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_Array()
	return ret
}

func NSMutableArrayWithCapacity(numItems NSUInteger) NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_ArrayWithCapacity((C.NSUInteger)(numItems))
	return ret
}

func NSMutableArrayAutomaticallyNotifiesObserversForKey(key NSString) bool {
	ret := (C.NSMutableArray_AutomaticallyNotifiesObserversForKey(key.Ptr())) != 0
	return ret
}

func NSMutableArrayCancelPreviousPerformRequestsWithTarget(aTarget NSObject)  {
	C.NSMutableArray_CancelPreviousPerformRequestsWithTarget(aTarget.Ptr())
}

func NSMutableArrayCancelPreviousPerformRequestsWithTargetSelector(aTarget NSObject, aSelector SEL, anArgument NSObject)  {
	C.NSMutableArray_CancelPreviousPerformRequestsWithTargetSelector(aTarget.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr())
}

func NSMutableArrayMutableCopyWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSMutableArray_MutableCopyWithZone(unsafe.Pointer(zone))
	return ret
}

func NSMutableArraySuperclass() Class {
	ret := (Class)(unsafe.Pointer(C.NSMutableArray_Superclass()))
	return ret
}

func NSMutableArrayInstanceMethodSignatureForSelector(aSelector SEL) NSMethodSignature {
	ret := NSMethodSignature{}
	ret.ptr = C.NSMutableArray_InstanceMethodSignatureForSelector(unsafe.Pointer(aSelector))
	return ret
}

func NSMutableArrayResolveClassMethod(sel SEL) bool {
	ret := (C.NSMutableArray_ResolveClassMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSMutableArrayWithObject(anObject NSObject) NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_ArrayWithObject(anObject.Ptr())
	return ret
}

func NSMutableArrayLoad()  {
	C.NSMutableArray_Load()
}

func NSMutableArrayWithContentsOfURL(url NSURL) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSMutableArray_ArrayWithContentsOfURL(url.Ptr())
	return ret
}

func NSMutableArrayWithContentsOfURLError(url NSURL, error *[]NSError) NSArray {

	goSlice1 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice1[i] = (*error)[i].Ptr()
	}
	ret := NSArray{}
	ret.ptr = C.NSMutableArray_ArrayWithContentsOfURLError(url.Ptr(), unsafe.Pointer(&goSlice1[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice1[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice1[i]
	}
	return ret
}

func NSMutableArrayIsSubclassOfClass(aClass Class) bool {
	ret := (C.NSMutableArray_IsSubclassOfClass(unsafe.Pointer(aClass))) != 0
	return ret
}

func NSMutableArrayVersion() NSInteger {
	ret := (NSInteger)(C.NSMutableArray_Version())
	return ret
}

func NSMutableArrayResolveInstanceMethod(sel SEL) bool {
	ret := (C.NSMutableArray_ResolveInstanceMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSMutableArrayClassForKeyedUnarchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSMutableArray_ClassForKeyedUnarchiver()))
	return ret
}

func NSMutableArrayNew() NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_New()
	return ret
}

func NSMutableArrayConformsToProtocol(protocol Protocol) bool {
	ret := (C.NSMutableArray_ConformsToProtocol(protocol.Ptr())) != 0
	return ret
}

func NSMutableArrayHash() NSUInteger {
	ret := (NSUInteger)(C.NSMutableArray_Hash())
	return ret
}

func NSMutableArrayWithContentsOfFile(path NSString) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSMutableArray_ArrayWithContentsOfFile(path.Ptr())
	return ret
}

func NSMutableArrayClass() Class {
	ret := (Class)(unsafe.Pointer(C.NSMutableArray_Class()))
	return ret
}

func NSMutableArrayCopyWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSMutableArray_CopyWithZone(unsafe.Pointer(zone))
	return ret
}

func NSMutableArraySetVersion(aVersion NSInteger)  {
	C.NSMutableArray_SetVersion((C.NSInteger)(aVersion))
}

func NSMutableArrayDebugDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSMutableArray_DebugDescription()
	return ret
}

func NSMutableArrayAllocWithZone(zone *_NSZone) NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_AllocWithZone(unsafe.Pointer(zone))
	return ret
}

func NSMutableArrayAccessInstanceVariablesDirectly() bool {
	ret := (C.NSMutableArray_AccessInstanceVariablesDirectly()) != 0
	return ret
}

func NSMutableArrayAlloc() NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_Alloc()
	return ret
}

func NSMutableArrayInstancesRespondToSelector(aSelector SEL) bool {
	ret := (C.NSMutableArray_InstancesRespondToSelector(unsafe.Pointer(aSelector))) != 0
	return ret
}

func NSMutableArrayWithObjectsCount(objects *[]Id, cnt NSUInteger) NSMutableArray {

	goSlice0 := make([]unsafe.Pointer,cap(*objects))
	for i := 0; i < len(*objects); i++ {
		goSlice0[i] = (*objects)[i].Ptr()
	}
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_ArrayWithObjectsCount(unsafe.Pointer(&goSlice0[0]), (C.NSUInteger)(cnt))
	(*objects) = (*objects)[:cap(*objects)]
	for i := 0; i < len(*objects); i++ {
		if goSlice0[i] == nil {
			(*objects) = (*objects)[:i]
			break
		}
		(*objects)[i].ptr = goSlice0[i]
	}
	return ret
}

func NSMutableArrayWithObjects(firstObj NSObject, objects ...NSObject) NSMutableArray {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_ArrayWithObjects(firstObj.Ptr(), unsafe.Pointer(&object))
	return ret
}

func (o NSMutableArray) RemoveObject(anObject NSObject)  {
	C.NSMutableArray_inst_RemoveObject(o.Ptr(), anObject.Ptr())
}

func (o NSMutableArray) RemoveObjectInRange(anObject NSObject, range_ NSRange)  {
	C.NSMutableArray_inst_RemoveObjectInRange(o.Ptr(), anObject.Ptr(), (C.NSRange)(range_))
}

func (o NSMutableArray) RemoveLastObject()  {
	C.NSMutableArray_inst_RemoveLastObject(o.Ptr())
}

func (o NSMutableArray) RemoveObjectsInArray(otherArray NSArray)  {
	C.NSMutableArray_inst_RemoveObjectsInArray(o.Ptr(), otherArray.Ptr())
}

func (o NSMutableArray) InitWithCoder(aDecoder NSCoder) NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_inst_InitWithCoder(o.Ptr(), aDecoder.Ptr())
	return ret
}

func (o NSMutableArray) AddObject(anObject NSObject)  {
	C.NSMutableArray_inst_AddObject(o.Ptr(), anObject.Ptr())
}

func (o NSMutableArray) RemoveObjectIdenticalTo(anObject NSObject)  {
	C.NSMutableArray_inst_RemoveObjectIdenticalTo(o.Ptr(), anObject.Ptr())
}

func (o NSMutableArray) RemoveObjectIdenticalToInRange(anObject NSObject, range_ NSRange)  {
	C.NSMutableArray_inst_RemoveObjectIdenticalToInRange(o.Ptr(), anObject.Ptr(), (C.NSRange)(range_))
}

func (o NSMutableArray) ExchangeObjectAtIndex(idx1 NSUInteger, idx2 NSUInteger)  {
	C.NSMutableArray_inst_ExchangeObjectAtIndex(o.Ptr(), (C.NSUInteger)(idx1), (C.NSUInteger)(idx2))
}

func (o NSMutableArray) RemoveObjectsAtIndexes(indexes NSIndexSet)  {
	C.NSMutableArray_inst_RemoveObjectsAtIndexes(o.Ptr(), indexes.Ptr())
}

func (o NSMutableArray) Init() NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_inst_Init(o.Ptr())
	return ret
}

func (o NSMutableArray) ReplaceObjectsInRangeWithObjectsFromArray(range_ NSRange, otherArray NSArray)  {
	C.NSMutableArray_inst_ReplaceObjectsInRangeWithObjectsFromArray(o.Ptr(), (C.NSRange)(range_), otherArray.Ptr())
}

func (o NSMutableArray) ReplaceObjectsInRangeWithObjectsFromArrayRange(range_ NSRange, otherArray NSArray, otherRange NSRange)  {
	C.NSMutableArray_inst_ReplaceObjectsInRangeWithObjectsFromArrayRange(o.Ptr(), (C.NSRange)(range_), otherArray.Ptr(), (C.NSRange)(otherRange))
}

func (o NSMutableArray) SetObject(obj NSObject, idx NSUInteger)  {
	C.NSMutableArray_inst_SetObject(o.Ptr(), obj.Ptr(), (C.NSUInteger)(idx))
}

func (o NSMutableArray) InsertObject(anObject NSObject, index NSUInteger)  {
	C.NSMutableArray_inst_InsertObject(o.Ptr(), anObject.Ptr(), (C.NSUInteger)(index))
}

func (o NSMutableArray) InsertObjects(objects NSArray, indexes NSIndexSet)  {
	C.NSMutableArray_inst_InsertObjects(o.Ptr(), objects.Ptr(), indexes.Ptr())
}

func (o NSMutableArray) SortUsingDescriptors(sortDescriptors NSArray)  {
	C.NSMutableArray_inst_SortUsingDescriptors(o.Ptr(), sortDescriptors.Ptr())
}

func (o NSMutableArray) RemoveAllObjects()  {
	C.NSMutableArray_inst_RemoveAllObjects(o.Ptr())
}

func (o NSMutableArray) SetArray(otherArray NSArray)  {
	C.NSMutableArray_inst_SetArray(o.Ptr(), otherArray.Ptr())
}

func (o NSMutableArray) RemoveObjectsInRange(range_ NSRange)  {
	C.NSMutableArray_inst_RemoveObjectsInRange(o.Ptr(), (C.NSRange)(range_))
}

func (o NSMutableArray) ReplaceObjectAtIndex(index NSUInteger, anObject NSObject)  {
	C.NSMutableArray_inst_ReplaceObjectAtIndex(o.Ptr(), (C.NSUInteger)(index), anObject.Ptr())
}

func (o NSMutableArray) SortUsingSelector(comparator SEL)  {
	C.NSMutableArray_inst_SortUsingSelector(o.Ptr(), unsafe.Pointer(comparator))
}

func (o NSMutableArray) InitWithContentsOfFile(path NSString) NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_inst_InitWithContentsOfFile(o.Ptr(), path.Ptr())
	return ret
}

func (o NSMutableArray) FilterUsingPredicate(predicate NSPredicate)  {
	C.NSMutableArray_inst_FilterUsingPredicate(o.Ptr(), predicate.Ptr())
}

func (o NSMutableArray) InitWithCapacity(numItems NSUInteger) NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_inst_InitWithCapacity(o.Ptr(), (C.NSUInteger)(numItems))
	return ret
}

func (o NSMutableArray) ReplaceObjectsAtIndexes(indexes NSIndexSet, objects NSArray)  {
	C.NSMutableArray_inst_ReplaceObjectsAtIndexes(o.Ptr(), indexes.Ptr(), objects.Ptr())
}

func (o NSMutableArray) InitWithContentsOfURL(url NSURL) NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSMutableArray_inst_InitWithContentsOfURL(o.Ptr(), url.Ptr())
	return ret
}

func (o NSMutableArray) AddObjectsFromArray(otherArray NSArray)  {
	C.NSMutableArray_inst_AddObjectsFromArray(o.Ptr(), otherArray.Ptr())
}

func (o NSMutableArray) RemoveObjectAtIndex(index NSUInteger)  {
	C.NSMutableArray_inst_RemoveObjectAtIndex(o.Ptr(), (C.NSUInteger)(index))
}

func NSStringPathWithComponents(components NSArray) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_PathWithComponents(components.Ptr())
	return ret
}

func NSStringDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_Description()
	return ret
}

func NSStringAccessInstanceVariablesDirectly() bool {
	ret := (C.NSString_AccessInstanceVariablesDirectly()) != 0
	return ret
}

func NSStringHash() NSUInteger {
	ret := (NSUInteger)(C.NSString_Hash())
	return ret
}

func NSStringWithString(string NSString) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_StringWithString(string.Ptr())
	return ret
}

func NSStringWithGoString(string string) NSString {
	string_chr := CharWithGoString(string)
	defer string_chr.Free()
	return NSStringWithString(NSStringWithUTF8String(string_chr))
}

func NSStringWithContentsOfURLEncoding(url NSURL, enc NSStringEncoding, error *[]NSError) NSString {

	goSlice2 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice2[i] = (*error)[i].Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_StringWithContentsOfURLEncoding(url.Ptr(), (C.NSStringEncoding)(enc), unsafe.Pointer(&goSlice2[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice2[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice2[i]
	}
	return ret
}

func NSStringWithContentsOfURLUsedEncoding(url NSURL, enc *NSStringEncoding, error *[]NSError) NSString {

	goSlice2 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice2[i] = (*error)[i].Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_StringWithContentsOfURLUsedEncoding(url.Ptr(), unsafe.Pointer(enc), unsafe.Pointer(&goSlice2[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice2[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice2[i]
	}
	return ret
}

func NSStringWithCString(cString *Char, enc NSStringEncoding) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_StringWithCString(unsafe.Pointer(cString), (C.NSStringEncoding)(enc))
	return ret
}

func NSStringInstancesRespondToSelector(aSelector SEL) bool {
	ret := (C.NSString_InstancesRespondToSelector(unsafe.Pointer(aSelector))) != 0
	return ret
}

func NSStringClassForKeyedUnarchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_ClassForKeyedUnarchiver()))
	return ret
}

func NSStringClass() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_Class()))
	return ret
}

func NSStringResolveClassMethod(sel SEL) bool {
	ret := (C.NSString_ResolveClassMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSStringWithContentsOfFileEncoding(path NSString, enc NSStringEncoding, error *[]NSError) NSString {

	goSlice2 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice2[i] = (*error)[i].Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_StringWithContentsOfFileEncoding(path.Ptr(), (C.NSStringEncoding)(enc), unsafe.Pointer(&goSlice2[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice2[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice2[i]
	}
	return ret
}

func NSStringWithContentsOfFileUsedEncoding(path NSString, enc *NSStringEncoding, error *[]NSError) NSString {

	goSlice2 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice2[i] = (*error)[i].Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_StringWithContentsOfFileUsedEncoding(path.Ptr(), unsafe.Pointer(enc), unsafe.Pointer(&goSlice2[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice2[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice2[i]
	}
	return ret
}

func NSStringDebugDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_DebugDescription()
	return ret
}

func NSStringConformsToProtocol(protocol Protocol) bool {
	ret := (C.NSString_ConformsToProtocol(protocol.Ptr())) != 0
	return ret
}

func NSStringDefaultCStringEncoding() NSStringEncoding {
	ret := (NSStringEncoding)(C.NSString_DefaultCStringEncoding())
	return ret
}

func NSStringMutableCopyWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSString_MutableCopyWithZone(unsafe.Pointer(zone))
	return ret
}

func NSStringCopyWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSString_CopyWithZone(unsafe.Pointer(zone))
	return ret
}

func NSStringLocalizedNameOfStringEncoding(encoding NSStringEncoding) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_LocalizedNameOfStringEncoding((C.NSStringEncoding)(encoding))
	return ret
}

func NSStringCancelPreviousPerformRequestsWithTarget(aTarget NSObject)  {
	C.NSString_CancelPreviousPerformRequestsWithTarget(aTarget.Ptr())
}

func NSStringCancelPreviousPerformRequestsWithTargetSelector(aTarget NSObject, aSelector SEL, anArgument NSObject)  {
	C.NSString_CancelPreviousPerformRequestsWithTargetSelector(aTarget.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr())
}

func NSStringVersion() NSInteger {
	ret := (NSInteger)(C.NSString_Version())
	return ret
}

func NSStringAutomaticallyNotifiesObserversForKey(key NSString) bool {
	ret := (C.NSString_AutomaticallyNotifiesObserversForKey(key.Ptr())) != 0
	return ret
}

func NSStringAvailableStringEncodings() *NSStringEncoding {
	ret := (*NSStringEncoding)(unsafe.Pointer(C.NSString_AvailableStringEncodings()))
	return ret
}

func NSStringIsSubclassOfClass(aClass Class) bool {
	ret := (C.NSString_IsSubclassOfClass(unsafe.Pointer(aClass))) != 0
	return ret
}

func NSStringSuperclass() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_Superclass()))
	return ret
}

func NSStringNew() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_New()
	return ret
}

func NSStringSetVersion(aVersion NSInteger)  {
	C.NSString_SetVersion((C.NSInteger)(aVersion))
}

func NSStringAllocWithZone(zone *_NSZone) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_AllocWithZone(unsafe.Pointer(zone))
	return ret
}

func NSStringClassFallbacksForKeyedArchiver() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSString_ClassFallbacksForKeyedArchiver()
	return ret
}

func NSStringLoad()  {
	C.NSString_Load()
}

func NSStringLocalizedStringWithFormat(format NSString, objects ...NSObject) NSString {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_LocalizedStringWithFormat(format.Ptr(), unsafe.Pointer(&object))
	return ret
}

func NSStringAlloc() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_Alloc()
	return ret
}

func NSStringWithCharacters(characters *Unichar, length NSUInteger) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_StringWithCharacters(unsafe.Pointer(characters), (C.NSUInteger)(length))
	return ret
}

func NSStringEncodingForData(data NSData, opts NSDictionary, string *[]NSString, usedLossyConversion *BOOL) NSStringEncoding {

	goSlice2 := make([]unsafe.Pointer,cap(*string))
	for i := 0; i < len(*string); i++ {
		goSlice2[i] = (*string)[i].Ptr()
	}
	ret := (NSStringEncoding)(C.NSString_StringEncodingForData(data.Ptr(), opts.Ptr(), unsafe.Pointer(&goSlice2[0]), unsafe.Pointer(usedLossyConversion)))
	(*string) = (*string)[:cap(*string)]
	for i := 0; i < len(*string); i++ {
		if goSlice2[i] == nil {
			(*string) = (*string)[:i]
			break
		}
		(*string)[i].ptr = goSlice2[i]
	}
	return ret
}

func NSStringResolveInstanceMethod(sel SEL) bool {
	ret := (C.NSString_ResolveInstanceMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSStringWithUTF8String(nullTerminatedCString *Char) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_StringWithUTF8String(unsafe.Pointer(nullTerminatedCString))
	return ret
}

func NSStringKeyPathsForValuesAffectingValueForKey(key NSString) NSSet {
	ret := NSSet{}
	ret.ptr = C.NSString_KeyPathsForValuesAffectingValueForKey(key.Ptr())
	return ret
}

func NSStringString() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_String()
	return ret
}

func NSStringInstanceMethodSignatureForSelector(aSelector SEL) NSMethodSignature {
	ret := NSMethodSignature{}
	ret.ptr = C.NSString_InstanceMethodSignatureForSelector(unsafe.Pointer(aSelector))
	return ret
}

func NSStringWithFormat(format NSString, objects ...NSObject) NSString {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_StringWithFormat(format.Ptr(), unsafe.Pointer(&object))
	return ret
}

func (o NSString) InitWithCharacters(characters *Unichar, length NSUInteger) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithCharacters(o.Ptr(), unsafe.Pointer(characters), (C.NSUInteger)(length))
	return ret
}

func (o NSString) StringByStandardizingPath() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByStandardizingPath(o.Ptr())
	return ret
}

func (o NSString) HasSuffix(str NSString) bool {
	ret := (C.NSString_inst_HasSuffix(o.Ptr(), str.Ptr())) != 0
	return ret
}

func (o NSString) ComponentsSeparatedByCharactersInSet(separator NSCharacterSet) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSString_inst_ComponentsSeparatedByCharactersInSet(o.Ptr(), separator.Ptr())
	return ret
}

func (o NSString) LengthOfBytesUsingEncoding(enc NSStringEncoding) NSUInteger {
	ret := (NSUInteger)(C.NSString_inst_LengthOfBytesUsingEncoding(o.Ptr(), (C.NSStringEncoding)(enc)))
	return ret
}

func (o NSString) LocalizedCaseInsensitiveContainsString(str NSString) bool {
	ret := (C.NSString_inst_LocalizedCaseInsensitiveContainsString(o.Ptr(), str.Ptr())) != 0
	return ret
}

func (o NSString) CapitalizedString() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_CapitalizedString(o.Ptr())
	return ret
}

func (o NSString) Hash() NSUInteger {
	ret := (NSUInteger)(C.NSString_inst_Hash(o.Ptr()))
	return ret
}

func (o NSString) InitWithContentsOfURLEncoding(url NSURL, enc NSStringEncoding, error *[]NSError) NSString {

	goSlice3 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice3[i] = (*error)[i].Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithContentsOfURLEncoding(o.Ptr(), url.Ptr(), (C.NSStringEncoding)(enc), unsafe.Pointer(&goSlice3[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice3[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice3[i]
	}
	return ret
}

func (o NSString) InitWithContentsOfURLUsedEncoding(url NSURL, enc *NSStringEncoding, error *[]NSError) NSString {

	goSlice3 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice3[i] = (*error)[i].Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithContentsOfURLUsedEncoding(o.Ptr(), url.Ptr(), unsafe.Pointer(enc), unsafe.Pointer(&goSlice3[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice3[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice3[i]
	}
	return ret
}

func (o NSString) CompletePathIntoString(outputName *[]NSString, flag BOOL, outputArray *[]NSArray, filterTypes NSArray) NSUInteger {

	goSlice1 := make([]unsafe.Pointer,cap(*outputName))
	for i := 0; i < len(*outputName); i++ {
		goSlice1[i] = (*outputName)[i].Ptr()
	}

	goSlice3 := make([]unsafe.Pointer,cap(*outputArray))
	for i := 0; i < len(*outputArray); i++ {
		goSlice3[i] = (*outputArray)[i].Ptr()
	}
	ret := (NSUInteger)(C.NSString_inst_CompletePathIntoString(o.Ptr(), unsafe.Pointer(&goSlice1[0]), (C.BOOL)(flag), unsafe.Pointer(&goSlice3[0]), filterTypes.Ptr()))
	(*outputName) = (*outputName)[:cap(*outputName)]
	for i := 0; i < len(*outputName); i++ {
		if goSlice1[i] == nil {
			(*outputName) = (*outputName)[:i]
			break
		}
		(*outputName)[i].ptr = goSlice1[i]
	}
	(*outputArray) = (*outputArray)[:cap(*outputArray)]
	for i := 0; i < len(*outputArray); i++ {
		if goSlice3[i] == nil {
			(*outputArray) = (*outputArray)[:i]
			break
		}
		(*outputArray)[i].ptr = goSlice3[i]
	}
	return ret
}

func (o NSString) DecomposedStringWithCanonicalMapping() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_DecomposedStringWithCanonicalMapping(o.Ptr())
	return ret
}

func (o NSString) InitWithFormat(format NSString, objects ...NSObject) NSString {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithFormat(o.Ptr(), format.Ptr(), unsafe.Pointer(&object))
	return ret
}

func (o NSString) InitWithFormatLocale(format NSString, locale NSObject, objects ...NSObject) NSString {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithFormatLocale(o.Ptr(), format.Ptr(), locale.Ptr(), unsafe.Pointer(&object))
	return ret
}

func (o NSString) SubstringWithRange(range_ NSRange) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_SubstringWithRange(o.Ptr(), (C.NSRange)(range_))
	return ret
}

func (o NSString) GetCharacters(buffer *Unichar)  {
	C.NSString_inst_GetCharacters(o.Ptr(), unsafe.Pointer(buffer))
}

func (o NSString) GetCharactersRange(buffer *Unichar, range_ NSRange)  {
	C.NSString_inst_GetCharactersRange(o.Ptr(), unsafe.Pointer(buffer), (C.NSRange)(range_))
}

func (o NSString) LowercaseStringWithLocale(locale NSLocale) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_LowercaseStringWithLocale(o.Ptr(), locale.Ptr())
	return ret
}

func (o NSString) StringByExpandingTildeInPath() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByExpandingTildeInPath(o.Ptr())
	return ret
}

func (o NSString) StringByAppendingFormat(format NSString, objects ...NSObject) NSString {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByAppendingFormat(o.Ptr(), format.Ptr(), unsafe.Pointer(&object))
	return ret
}

func (o NSString) RangeOfComposedCharacterSequencesForRange(range_ NSRange) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfComposedCharacterSequencesForRange(o.Ptr(), (C.NSRange)(range_)))
	return ret
}

func (o NSString) PropertyListFromStringsFileFormat() NSDictionary {
	ret := NSDictionary{}
	ret.ptr = C.NSString_inst_PropertyListFromStringsFileFormat(o.Ptr())
	return ret
}

func (o NSString) CommonPrefixWithString(str NSString, mask NSStringCompareOptions) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_CommonPrefixWithString(o.Ptr(), str.Ptr(), (C.NSStringCompareOptions)(mask))
	return ret
}

func (o NSString) CommonPrefixWithGoString(str string, mask NSStringCompareOptions) NSString {
	str_chr := CharWithGoString(str)
	defer str_chr.Free()
	return o.CommonPrefixWithString(NSStringWithUTF8String(str_chr), mask)
}

func (o NSString) FileSystemRepresentation() *Char {
	ret := (*Char)(unsafe.Pointer(C.NSString_inst_FileSystemRepresentation(o.Ptr())))
	return ret
}

func (o NSString) StringByFoldingWithOptions(options NSStringCompareOptions, locale NSLocale) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByFoldingWithOptions(o.Ptr(), (C.NSStringCompareOptions)(options), locale.Ptr())
	return ret
}

func (o NSString) StringsByAppendingPaths(paths NSArray) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSString_inst_StringsByAppendingPaths(o.Ptr(), paths.Ptr())
	return ret
}

func (o NSString) InitWithCharactersNoCopy(characters *Unichar, length NSUInteger, freeBuffer BOOL) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithCharactersNoCopy(o.Ptr(), unsafe.Pointer(characters), (C.NSUInteger)(length), (C.BOOL)(freeBuffer))
	return ret
}

func (o NSString) LocalizedStandardCompare(string NSString) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_LocalizedStandardCompare(o.Ptr(), string.Ptr()))
	return ret
}

func (o NSString) LocalizedCapitalizedString() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_LocalizedCapitalizedString(o.Ptr())
	return ret
}

func (o NSString) UppercaseString() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_UppercaseString(o.Ptr())
	return ret
}

func (o NSString) PropertyList() Id {
	ret := Id{}
	ret.ptr = C.NSString_inst_PropertyList(o.Ptr())
	return ret
}

func (o NSString) LocalizedStandardRangeOfString(str NSString) NSRange {
	ret := (NSRange)(C.NSString_inst_LocalizedStandardRangeOfString(o.Ptr(), str.Ptr()))
	return ret
}

func (o NSString) WriteToFile(path NSString, useAuxiliaryFile BOOL, enc NSStringEncoding, error *[]NSError) bool {

	goSlice4 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice4[i] = (*error)[i].Ptr()
	}
	ret := (C.NSString_inst_WriteToFile(o.Ptr(), path.Ptr(), (C.BOOL)(useAuxiliaryFile), (C.NSStringEncoding)(enc), unsafe.Pointer(&goSlice4[0]))) != 0
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice4[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice4[i]
	}
	return ret
}

func (o NSString) CharacterAtIndex(index NSUInteger) Unichar {
	ret := (Unichar)(C.NSString_inst_CharacterAtIndex(o.Ptr(), (C.NSUInteger)(index)))
	return ret
}

func (o NSString) StringByDeletingPathExtension() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByDeletingPathExtension(o.Ptr())
	return ret
}

func (o NSString) StringByTrimmingCharactersInSet(set NSCharacterSet) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByTrimmingCharactersInSet(o.Ptr(), set.Ptr())
	return ret
}

func (o NSString) PrecomposedStringWithCompatibilityMapping() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_PrecomposedStringWithCompatibilityMapping(o.Ptr())
	return ret
}

func (o NSString) ComponentsSeparatedByString(separator NSString) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSString_inst_ComponentsSeparatedByString(o.Ptr(), separator.Ptr())
	return ret
}

func (o NSString) StringByDeletingLastPathComponent() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByDeletingLastPathComponent(o.Ptr())
	return ret
}

func (o NSString) PrecomposedStringWithCanonicalMapping() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_PrecomposedStringWithCanonicalMapping(o.Ptr())
	return ret
}

func (o NSString) GetParagraphStart(startPtr *NSUInteger, parEndPtr *NSUInteger, contentsEndPtr *NSUInteger, range_ NSRange)  {
	C.NSString_inst_GetParagraphStart(o.Ptr(), unsafe.Pointer(startPtr), unsafe.Pointer(parEndPtr), unsafe.Pointer(contentsEndPtr), (C.NSRange)(range_))
}

func (o NSString) InitWithCoder(aDecoder NSCoder) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithCoder(o.Ptr(), aDecoder.Ptr())
	return ret
}

func (o NSString) LineRangeForRange(range_ NSRange) NSRange {
	ret := (NSRange)(C.NSString_inst_LineRangeForRange(o.Ptr(), (C.NSRange)(range_)))
	return ret
}

func (o NSString) RangeOfComposedCharacterSequenceAtIndex(index NSUInteger) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfComposedCharacterSequenceAtIndex(o.Ptr(), (C.NSUInteger)(index)))
	return ret
}

func (o NSString) StringByAppendingPathComponent(str NSString) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByAppendingPathComponent(o.Ptr(), str.Ptr())
	return ret
}

func (o NSString) WriteToURL(url NSURL, useAuxiliaryFile BOOL, enc NSStringEncoding, error *[]NSError) bool {

	goSlice4 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice4[i] = (*error)[i].Ptr()
	}
	ret := (C.NSString_inst_WriteToURL(o.Ptr(), url.Ptr(), (C.BOOL)(useAuxiliaryFile), (C.NSStringEncoding)(enc), unsafe.Pointer(&goSlice4[0]))) != 0
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice4[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice4[i]
	}
	return ret
}

func (o NSString) InitWithBytes(bytes unsafe.Pointer, len_ NSUInteger, encoding NSStringEncoding) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithBytes(o.Ptr(), unsafe.Pointer(bytes), (C.NSUInteger)(len_), (C.NSStringEncoding)(encoding))
	return ret
}

func (o NSString) DataUsingEncoding(encoding NSStringEncoding) NSData {
	ret := NSData{}
	ret.ptr = C.NSString_inst_DataUsingEncoding(o.Ptr(), (C.NSStringEncoding)(encoding))
	return ret
}

func (o NSString) DataUsingEncodingAllowLossyConversion(encoding NSStringEncoding, lossy BOOL) NSData {
	ret := NSData{}
	ret.ptr = C.NSString_inst_DataUsingEncodingAllowLossyConversion(o.Ptr(), (C.NSStringEncoding)(encoding), (C.BOOL)(lossy))
	return ret
}

func (o NSString) GetLineStart(startPtr *NSUInteger, lineEndPtr *NSUInteger, contentsEndPtr *NSUInteger, range_ NSRange)  {
	C.NSString_inst_GetLineStart(o.Ptr(), unsafe.Pointer(startPtr), unsafe.Pointer(lineEndPtr), unsafe.Pointer(contentsEndPtr), (C.NSRange)(range_))
}

func (o NSString) RangeOfCharacterFromSet(searchSet NSCharacterSet) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfCharacterFromSet(o.Ptr(), searchSet.Ptr()))
	return ret
}

func (o NSString) RangeOfCharacterFromSetOptions(searchSet NSCharacterSet, mask NSStringCompareOptions) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfCharacterFromSetOptions(o.Ptr(), searchSet.Ptr(), (C.NSStringCompareOptions)(mask)))
	return ret
}

func (o NSString) RangeOfCharacterFromSetOptionsRange(searchSet NSCharacterSet, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfCharacterFromSetOptionsRange(o.Ptr(), searchSet.Ptr(), (C.NSStringCompareOptions)(mask), (C.NSRange)(rangeOfReceiverToSearch)))
	return ret
}

func (o NSString) LastPathComponent() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_LastPathComponent(o.Ptr())
	return ret
}

func (o NSString) StringByResolvingSymlinksInPath() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByResolvingSymlinksInPath(o.Ptr())
	return ret
}

func (o NSString) PathExtension() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_PathExtension(o.Ptr())
	return ret
}

func (o NSString) Length() NSUInteger {
	ret := (NSUInteger)(C.NSString_inst_Length(o.Ptr()))
	return ret
}

func (o NSString) StringByReplacingCharactersInRange(range_ NSRange, replacement NSString) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByReplacingCharactersInRange(o.Ptr(), (C.NSRange)(range_), replacement.Ptr())
	return ret
}

func (o NSString) MaximumLengthOfBytesUsingEncoding(enc NSStringEncoding) NSUInteger {
	ret := (NSUInteger)(C.NSString_inst_MaximumLengthOfBytesUsingEncoding(o.Ptr(), (C.NSStringEncoding)(enc)))
	return ret
}

func (o NSString) UppercaseStringWithLocale(locale NSLocale) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_UppercaseStringWithLocale(o.Ptr(), locale.Ptr())
	return ret
}

func (o NSString) HasPrefix(str NSString) bool {
	ret := (C.NSString_inst_HasPrefix(o.Ptr(), str.Ptr())) != 0
	return ret
}

func (o NSString) InitWithUTF8String(nullTerminatedCString *Char) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithUTF8String(o.Ptr(), unsafe.Pointer(nullTerminatedCString))
	return ret
}

func (o NSString) VariantFittingPresentationWidth(width NSInteger) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_VariantFittingPresentationWidth(o.Ptr(), (C.NSInteger)(width))
	return ret
}

func (o NSString) SubstringFromIndex(from NSUInteger) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_SubstringFromIndex(o.Ptr(), (C.NSUInteger)(from))
	return ret
}

func (o NSString) SubstringToIndex(to NSUInteger) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_SubstringToIndex(o.Ptr(), (C.NSUInteger)(to))
	return ret
}

func (o NSString) Compare(string NSString) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_Compare(o.Ptr(), string.Ptr()))
	return ret
}

func (o NSString) CompareOptions(string NSString, mask NSStringCompareOptions) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_CompareOptions(o.Ptr(), string.Ptr(), (C.NSStringCompareOptions)(mask)))
	return ret
}

func (o NSString) CompareOptionsRange(string NSString, mask NSStringCompareOptions, rangeOfReceiverToCompare NSRange) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_CompareOptionsRange(o.Ptr(), string.Ptr(), (C.NSStringCompareOptions)(mask), (C.NSRange)(rangeOfReceiverToCompare)))
	return ret
}

func (o NSString) CompareOptionsRangeLocale(string NSString, mask NSStringCompareOptions, rangeOfReceiverToCompare NSRange, locale NSObject) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_CompareOptionsRangeLocale(o.Ptr(), string.Ptr(), (C.NSStringCompareOptions)(mask), (C.NSRange)(rangeOfReceiverToCompare), locale.Ptr()))
	return ret
}

func (o NSString) IntegerValue() NSInteger {
	ret := (NSInteger)(C.NSString_inst_IntegerValue(o.Ptr()))
	return ret
}

func (o NSString) StringByRemovingPercentEncoding() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByRemovingPercentEncoding(o.Ptr())
	return ret
}

func (o NSString) InitWithCString(nullTerminatedCString *Char, encoding NSStringEncoding) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithCString(o.Ptr(), unsafe.Pointer(nullTerminatedCString), (C.NSStringEncoding)(encoding))
	return ret
}

func (o NSString) IsEqualToString(aString NSString) bool {
	ret := (C.NSString_inst_IsEqualToString(o.Ptr(), aString.Ptr())) != 0
	return ret
}

func (o NSString) LocalizedLowercaseString() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_LocalizedLowercaseString(o.Ptr())
	return ret
}

func (o NSString) StringByAbbreviatingWithTildeInPath() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByAbbreviatingWithTildeInPath(o.Ptr())
	return ret
}

func (o NSString) StringByApplyingTransform(transform NSStringTransform, reverse BOOL) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByApplyingTransform(o.Ptr(), transform.Ptr(), (C.BOOL)(reverse))
	return ret
}

func (o NSString) DecomposedStringWithCompatibilityMapping() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_DecomposedStringWithCompatibilityMapping(o.Ptr())
	return ret
}

func (o NSString) LocalizedCaseInsensitiveCompare(string NSString) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_LocalizedCaseInsensitiveCompare(o.Ptr(), string.Ptr()))
	return ret
}

func (o NSString) FastestEncoding() NSStringEncoding {
	ret := (NSStringEncoding)(C.NSString_inst_FastestEncoding(o.Ptr()))
	return ret
}

func (o NSString) CapitalizedStringWithLocale(locale NSLocale) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_CapitalizedStringWithLocale(o.Ptr(), locale.Ptr())
	return ret
}

func (o NSString) ContainsString(str NSString) bool {
	ret := (C.NSString_inst_ContainsString(o.Ptr(), str.Ptr())) != 0
	return ret
}

func (o NSString) InitWithString(aString NSString) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithString(o.Ptr(), aString.Ptr())
	return ret
}

func (o NSString) InitWithGoString(aString string) NSString {
	aString_chr := CharWithGoString(aString)
	defer aString_chr.Free()
	return o.InitWithString(NSStringWithUTF8String(aString_chr))
}

func (o NSString) BoolValue() bool {
	ret := (C.NSString_inst_BoolValue(o.Ptr())) != 0
	return ret
}

func (o NSString) SmallestEncoding() NSStringEncoding {
	ret := (NSStringEncoding)(C.NSString_inst_SmallestEncoding(o.Ptr()))
	return ret
}

func (o NSString) LocalizedCompare(string NSString) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_LocalizedCompare(o.Ptr(), string.Ptr()))
	return ret
}

func (o NSString) CStringUsingEncoding(encoding NSStringEncoding) *Char {
	ret := (*Char)(unsafe.Pointer(C.NSString_inst_CStringUsingEncoding(o.Ptr(), (C.NSStringEncoding)(encoding))))
	return ret
}

func (o NSString) LocalizedStandardContainsString(str NSString) bool {
	ret := (C.NSString_inst_LocalizedStandardContainsString(o.Ptr(), str.Ptr())) != 0
	return ret
}

func (o NSString) StringByAppendingPathExtension(str NSString) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByAppendingPathExtension(o.Ptr(), str.Ptr())
	return ret
}

func (o NSString) InitWithBytesNoCopy(bytes unsafe.Pointer, len_ NSUInteger, encoding NSStringEncoding, freeBuffer BOOL) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithBytesNoCopy(o.Ptr(), unsafe.Pointer(bytes), (C.NSUInteger)(len_), (C.NSStringEncoding)(encoding), (C.BOOL)(freeBuffer))
	return ret
}

func (o NSString) PathComponents() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSString_inst_PathComponents(o.Ptr())
	return ret
}

func (o NSString) DoubleValue() Double {
	ret := (Double)(C.NSString_inst_DoubleValue(o.Ptr()))
	return ret
}

func (o NSString) InitWithContentsOfFileEncoding(path NSString, enc NSStringEncoding, error *[]NSError) NSString {

	goSlice3 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice3[i] = (*error)[i].Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithContentsOfFileEncoding(o.Ptr(), path.Ptr(), (C.NSStringEncoding)(enc), unsafe.Pointer(&goSlice3[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice3[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice3[i]
	}
	return ret
}

func (o NSString) InitWithContentsOfFileUsedEncoding(path NSString, enc *NSStringEncoding, error *[]NSError) NSString {

	goSlice3 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice3[i] = (*error)[i].Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithContentsOfFileUsedEncoding(o.Ptr(), path.Ptr(), unsafe.Pointer(enc), unsafe.Pointer(&goSlice3[0]))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice3[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		(*error)[i].ptr = goSlice3[i]
	}
	return ret
}

func (o NSString) GetCString(buffer *Char, maxBufferCount NSUInteger, encoding NSStringEncoding) bool {
	ret := (C.NSString_inst_GetCString(o.Ptr(), unsafe.Pointer(buffer), (C.NSUInteger)(maxBufferCount), (C.NSStringEncoding)(encoding))) != 0
	return ret
}

func (o NSString) StringByAppendingString(aString NSString) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByAppendingString(o.Ptr(), aString.Ptr())
	return ret
}

func (o NSString) GetBytes(buffer unsafe.Pointer, maxBufferCount NSUInteger, usedBufferCount *NSUInteger, encoding NSStringEncoding, options NSStringEncodingConversionOptions, range_ NSRange, leftover NSRangePointer) bool {
	ret := (C.NSString_inst_GetBytes(o.Ptr(), unsafe.Pointer(buffer), (C.NSUInteger)(maxBufferCount), unsafe.Pointer(usedBufferCount), (C.NSStringEncoding)(encoding), (C.NSStringEncodingConversionOptions)(options), (C.NSRange)(range_), unsafe.Pointer(leftover))) != 0
	return ret
}

func (o NSString) Init() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_Init(o.Ptr())
	return ret
}

func (o NSString) InitWithData(data NSData, encoding NSStringEncoding) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_InitWithData(o.Ptr(), data.Ptr(), (C.NSStringEncoding)(encoding))
	return ret
}

func (o NSString) CaseInsensitiveCompare(string NSString) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_CaseInsensitiveCompare(o.Ptr(), string.Ptr()))
	return ret
}

func (o NSString) GetFileSystemRepresentation(cname *Char, max NSUInteger) bool {
	ret := (C.NSString_inst_GetFileSystemRepresentation(o.Ptr(), unsafe.Pointer(cname), (C.NSUInteger)(max))) != 0
	return ret
}

func (o NSString) LowercaseString() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_LowercaseString(o.Ptr())
	return ret
}

func (o NSString) StringByPaddingToLength(newLength NSUInteger, padString NSString, padIndex NSUInteger) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByPaddingToLength(o.Ptr(), (C.NSUInteger)(newLength), padString.Ptr(), (C.NSUInteger)(padIndex))
	return ret
}

func (o NSString) StringByReplacingOccurrencesOfStringWithString(target NSString, replacement NSString) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByReplacingOccurrencesOfStringWithString(o.Ptr(), target.Ptr(), replacement.Ptr())
	return ret
}

func (o NSString) StringByReplacingOccurrencesOfStringWithStringOptions(target NSString, replacement NSString, options NSStringCompareOptions, searchRange NSRange) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByReplacingOccurrencesOfStringWithStringOptions(o.Ptr(), target.Ptr(), replacement.Ptr(), (C.NSStringCompareOptions)(options), (C.NSRange)(searchRange))
	return ret
}

func (o NSString) IntValue() Int {
	ret := (Int)(C.NSString_inst_IntValue(o.Ptr()))
	return ret
}

func (o NSString) CanBeConvertedToEncoding(encoding NSStringEncoding) bool {
	ret := (C.NSString_inst_CanBeConvertedToEncoding(o.Ptr(), (C.NSStringEncoding)(encoding))) != 0
	return ret
}

func (o NSString) StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters NSCharacterSet) NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_StringByAddingPercentEncodingWithAllowedCharacters(o.Ptr(), allowedCharacters.Ptr())
	return ret
}

func (o NSString) UTF8String() *Char {
	ret := (*Char)(unsafe.Pointer(C.NSString_inst_UTF8String(o.Ptr())))
	return ret
}

func (o NSString) ParagraphRangeForRange(range_ NSRange) NSRange {
	ret := (NSRange)(C.NSString_inst_ParagraphRangeForRange(o.Ptr(), (C.NSRange)(range_)))
	return ret
}

func (o NSString) LinguisticTagsInRange(range_ NSRange, scheme NSLinguisticTagScheme, options NSLinguisticTaggerOptions, orthography NSOrthography, tokenRanges *[]NSArray) NSArray {

	goSlice5 := make([]unsafe.Pointer,cap(*tokenRanges))
	for i := 0; i < len(*tokenRanges); i++ {
		goSlice5[i] = (*tokenRanges)[i].Ptr()
	}
	ret := NSArray{}
	ret.ptr = C.NSString_inst_LinguisticTagsInRange(o.Ptr(), (C.NSRange)(range_), scheme.Ptr(), (C.NSLinguisticTaggerOptions)(options), orthography.Ptr(), unsafe.Pointer(&goSlice5[0]))
	(*tokenRanges) = (*tokenRanges)[:cap(*tokenRanges)]
	for i := 0; i < len(*tokenRanges); i++ {
		if goSlice5[i] == nil {
			(*tokenRanges) = (*tokenRanges)[:i]
			break
		}
		(*tokenRanges)[i].ptr = goSlice5[i]
	}
	return ret
}

func (o NSString) Description() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_Description(o.Ptr())
	return ret
}

func (o NSString) FloatValue() Float {
	ret := (Float)(C.NSString_inst_FloatValue(o.Ptr()))
	return ret
}

func (o NSString) LocalizedUppercaseString() NSString {
	ret := NSString{}
	ret.ptr = C.NSString_inst_LocalizedUppercaseString(o.Ptr())
	return ret
}

func (o NSString) LongLongValue() LongLong {
	ret := (LongLong)(C.NSString_inst_LongLongValue(o.Ptr()))
	return ret
}

func (o NSString) IsAbsolutePath() bool {
	ret := (C.NSString_inst_IsAbsolutePath(o.Ptr())) != 0
	return ret
}

func (o NSString) RangeOfString(searchString NSString) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfString(o.Ptr(), searchString.Ptr()))
	return ret
}

func (o NSString) RangeOfStringOptions(searchString NSString, mask NSStringCompareOptions) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfStringOptions(o.Ptr(), searchString.Ptr(), (C.NSStringCompareOptions)(mask)))
	return ret
}

func (o NSString) RangeOfStringOptionsRange(searchString NSString, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfStringOptionsRange(o.Ptr(), searchString.Ptr(), (C.NSStringCompareOptions)(mask), (C.NSRange)(rangeOfReceiverToSearch)))
	return ret
}

func (o NSString) RangeOfStringOptionsRangeLocale(searchString NSString, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange, locale NSLocale) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfStringOptionsRangeLocale(o.Ptr(), searchString.Ptr(), (C.NSStringCompareOptions)(mask), (C.NSRange)(rangeOfReceiverToSearch), locale.Ptr()))
	return ret
}

func (o NSString) CopyWithZone(zone *NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSString_inst_CopyWithZone(o.Ptr(), unsafe.Pointer(zone))
	return ret
}

func (o NSString) MutableCopyWithZone(zone *NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSString_inst_MutableCopyWithZone(o.Ptr(), unsafe.Pointer(zone))
	return ret
}

func NSStringSupportsSecureCoding() bool {
	ret := (C.NSString_SupportsSecureCoding()) != 0
	return ret
}

func NSStringObjectWithItemProviderData(data NSData, typeIdentifier NSString, outError *[]NSError) NSString {

	goSlice2 := make([]unsafe.Pointer,cap(*outError))
	for i := 0; i < len(*outError); i++ {
		goSlice2[i] = (*outError)[i].Ptr()
	}
	ret := NSString{}
	ret.ptr = C.NSString_ObjectWithItemProviderData(data.Ptr(), typeIdentifier.Ptr(), unsafe.Pointer(&goSlice2[0]))
	(*outError) = (*outError)[:cap(*outError)]
	for i := 0; i < len(*outError); i++ {
		if goSlice2[i] == nil {
			(*outError) = (*outError)[:i]
			break
		}
		(*outError)[i].ptr = goSlice2[i]
	}
	return ret
}

func NSStringReadableTypeIdentifiersForItemProvider() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSString_ReadableTypeIdentifiersForItemProvider()
	return ret
}

func NSStringItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier NSString) NSItemProviderRepresentationVisibility {
	ret := (NSItemProviderRepresentationVisibility)(C.NSString_ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier.Ptr()))
	return ret
}

func NSStringWritableTypeIdentifiersForItemProvider() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSString_WritableTypeIdentifiersForItemProvider()
	return ret
}

func (o NSString) ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier NSString) NSItemProviderRepresentationVisibility {
	ret := (NSItemProviderRepresentationVisibility)(C.NSString_inst_ItemProviderVisibilityForRepresentationWithTypeIdentifier(o.Ptr(), typeIdentifier.Ptr()))
	return ret
}

func (o NSString) WritableTypeIdentifiersForItemProvider() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSString_inst_WritableTypeIdentifiersForItemProvider(o.Ptr())
	return ret
}

func NSObjectDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSObject_Description()
	return ret
}

func NSObjectCopyWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSObject_CopyWithZone(unsafe.Pointer(zone))
	return ret
}

func NSObjectMutableCopyWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSObject_MutableCopyWithZone(unsafe.Pointer(zone))
	return ret
}

func NSObjectCancelPreviousPerformRequestsWithTarget(aTarget NSObject)  {
	C.NSObject_CancelPreviousPerformRequestsWithTarget(aTarget.Ptr())
}

func NSObjectCancelPreviousPerformRequestsWithTargetSelector(aTarget NSObject, aSelector SEL, anArgument NSObject)  {
	C.NSObject_CancelPreviousPerformRequestsWithTargetSelector(aTarget.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr())
}

func NSObjectHash() NSUInteger {
	ret := (NSUInteger)(C.NSObject_Hash())
	return ret
}

func NSObjectAllocWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSObject_AllocWithZone(unsafe.Pointer(zone))
	return ret
}

func NSObjectDebugDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSObject_DebugDescription()
	return ret
}

func NSObjectIsSubclassOfClass(aClass Class) bool {
	ret := (C.NSObject_IsSubclassOfClass(unsafe.Pointer(aClass))) != 0
	return ret
}

func NSObjectInstancesRespondToSelector(aSelector SEL) bool {
	ret := (C.NSObject_InstancesRespondToSelector(unsafe.Pointer(aSelector))) != 0
	return ret
}

func NSObjectKeyPathsForValuesAffectingValueForKey(key NSString) NSSet {
	ret := NSSet{}
	ret.ptr = C.NSObject_KeyPathsForValuesAffectingValueForKey(key.Ptr())
	return ret
}

func NSObjectResolveClassMethod(sel SEL) bool {
	ret := (C.NSObject_ResolveClassMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSObjectAlloc() Id {
	ret := Id{}
	ret.ptr = C.NSObject_Alloc()
	return ret
}

func NSObjectClassFallbacksForKeyedArchiver() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSObject_ClassFallbacksForKeyedArchiver()
	return ret
}

func NSObjectClassForKeyedUnarchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_ClassForKeyedUnarchiver()))
	return ret
}

func NSObjectSetVersion(aVersion NSInteger)  {
	C.NSObject_SetVersion((C.NSInteger)(aVersion))
}

func NSObjectVersion() NSInteger {
	ret := (NSInteger)(C.NSObject_Version())
	return ret
}

func NSObjectConformsToProtocol(protocol Protocol) bool {
	ret := (C.NSObject_ConformsToProtocol(protocol.Ptr())) != 0
	return ret
}

func NSObjectInstanceMethodSignatureForSelector(aSelector SEL) NSMethodSignature {
	ret := NSMethodSignature{}
	ret.ptr = C.NSObject_InstanceMethodSignatureForSelector(unsafe.Pointer(aSelector))
	return ret
}

func NSObjectSuperclass() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_Superclass()))
	return ret
}

func NSObjectAutomaticallyNotifiesObserversForKey(key NSString) bool {
	ret := (C.NSObject_AutomaticallyNotifiesObserversForKey(key.Ptr())) != 0
	return ret
}

func NSObjectLoad()  {
	C.NSObject_Load()
}

func NSObjectAccessInstanceVariablesDirectly() bool {
	ret := (C.NSObject_AccessInstanceVariablesDirectly()) != 0
	return ret
}

func NSObjectNew() Id {
	ret := Id{}
	ret.ptr = C.NSObject_New()
	return ret
}

func NSObjectResolveInstanceMethod(sel SEL) bool {
	ret := (C.NSObject_ResolveInstanceMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSObjectClass() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_Class()))
	return ret
}

func (o Id) AutoContentAccessingProxy() Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_AutoContentAccessingProxy(o.Ptr())
	return ret
}

func (o Id) IsLessThanOrEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_IsLessThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) MutableCopy() Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_MutableCopy(o.Ptr())
	return ret
}

func (o Id) DoesNotRecognizeSelector(aSelector SEL)  {
	C.NSObject_inst_DoesNotRecognizeSelector(o.Ptr(), unsafe.Pointer(aSelector))
}

func (o Id) DictionaryWithValuesForKeys(keys NSArray) NSDictionary {
	ret := NSDictionary{}
	ret.ptr = C.NSObject_inst_DictionaryWithValuesForKeys(o.Ptr(), keys.Ptr())
	return ret
}

func (o Id) ObservationInfo() unsafe.Pointer {
	ret := (unsafe.Pointer)(unsafe.Pointer(C.NSObject_inst_ObservationInfo(o.Ptr())))
	return ret
}

func (o Id) AttemptRecoveryFromErrorOptionIndex(error NSError, recoveryOptionIndex NSUInteger) bool {
	ret := (C.NSObject_inst_AttemptRecoveryFromErrorOptionIndex(o.Ptr(), error.Ptr(), (C.NSUInteger)(recoveryOptionIndex))) != 0
	return ret
}

func (o Id) AttemptRecoveryFromErrorOptionIndexDelegate(error NSError, recoveryOptionIndex NSUInteger, delegate NSObject, didRecoverSelector SEL, contextInfo unsafe.Pointer)  {
	C.NSObject_inst_AttemptRecoveryFromErrorOptionIndexDelegate(o.Ptr(), error.Ptr(), (C.NSUInteger)(recoveryOptionIndex), delegate.Ptr(), unsafe.Pointer(didRecoverSelector), unsafe.Pointer(contextInfo))
}

func (o Id) InsertValueInPropertyWithKey(value NSObject, key NSString)  {
	C.NSObject_inst_InsertValueInPropertyWithKey(o.Ptr(), value.Ptr(), key.Ptr())
}

func (o Id) InsertValueAtIndex(value NSObject, index NSUInteger, key NSString)  {
	C.NSObject_inst_InsertValueAtIndex(o.Ptr(), value.Ptr(), (C.NSUInteger)(index), key.Ptr())
}

func (o Id) WillChange(changeKind NSKeyValueChange, indexes NSIndexSet, key NSString)  {
	C.NSObject_inst_WillChange(o.Ptr(), (C.NSKeyValueChange)(changeKind), indexes.Ptr(), key.Ptr())
}

func (o Id) MutableSetValueForKeyPath(keyPath NSString) NSMutableSet {
	ret := NSMutableSet{}
	ret.ptr = C.NSObject_inst_MutableSetValueForKeyPath(o.Ptr(), keyPath.Ptr())
	return ret
}

func (o Id) ObserveValueForKeyPath(keyPath NSString, object NSObject, change NSDictionary, context unsafe.Pointer)  {
	C.NSObject_inst_ObserveValueForKeyPath(o.Ptr(), keyPath.Ptr(), object.Ptr(), change.Ptr(), unsafe.Pointer(context))
}

func (o Id) ScriptingBeginsWith(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingBeginsWith(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) MutableArrayValueForKeyPath(keyPath NSString) NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSObject_inst_MutableArrayValueForKeyPath(o.Ptr(), keyPath.Ptr())
	return ret
}

func (o Id) AttributeKeys() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSObject_inst_AttributeKeys(o.Ptr())
	return ret
}

func (o Id) SetValueForKey(value NSObject, key NSString)  {
	C.NSObject_inst_SetValueForKey(o.Ptr(), value.Ptr(), key.Ptr())
}

func (o Id) SetValueForKeyPath(value NSObject, keyPath NSString)  {
	C.NSObject_inst_SetValueForKeyPath(o.Ptr(), value.Ptr(), keyPath.Ptr())
}

func (o Id) SetValueForUndefinedKey(value NSObject, key NSString)  {
	C.NSObject_inst_SetValueForUndefinedKey(o.Ptr(), value.Ptr(), key.Ptr())
}

func (o Id) ValueForKeyPath(keyPath NSString) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_ValueForKeyPath(o.Ptr(), keyPath.Ptr())
	return ret
}

func (o Id) CoerceValue(value NSObject, key NSString) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_CoerceValue(o.Ptr(), value.Ptr(), key.Ptr())
	return ret
}

func (o Id) DidChangeValueForKey(key NSString)  {
	C.NSObject_inst_DidChangeValueForKey(o.Ptr(), key.Ptr())
}

func (o Id) DidChangeValueForKeyWithSetMutation(key NSString, mutationKind NSKeyValueSetMutationKind, objects NSSet)  {
	C.NSObject_inst_DidChangeValueForKeyWithSetMutation(o.Ptr(), key.Ptr(), (C.NSKeyValueSetMutationKind)(mutationKind), objects.Ptr())
}

func (o Id) ValidateValueForKey(ioValue *[]Id, inKey NSString, outError *[]NSError) bool {

	goSlice1 := make([]unsafe.Pointer,cap(*ioValue))
	for i := 0; i < len(*ioValue); i++ {
		goSlice1[i] = (*ioValue)[i].Ptr()
	}

	goSlice3 := make([]unsafe.Pointer,cap(*outError))
	for i := 0; i < len(*outError); i++ {
		goSlice3[i] = (*outError)[i].Ptr()
	}
	ret := (C.NSObject_inst_ValidateValueForKey(o.Ptr(), unsafe.Pointer(&goSlice1[0]), inKey.Ptr(), unsafe.Pointer(&goSlice3[0]))) != 0
	(*ioValue) = (*ioValue)[:cap(*ioValue)]
	for i := 0; i < len(*ioValue); i++ {
		if goSlice1[i] == nil {
			(*ioValue) = (*ioValue)[:i]
			break
		}
		(*ioValue)[i].ptr = goSlice1[i]
	}
	(*outError) = (*outError)[:cap(*outError)]
	for i := 0; i < len(*outError); i++ {
		if goSlice3[i] == nil {
			(*outError) = (*outError)[:i]
			break
		}
		(*outError)[i].ptr = goSlice3[i]
	}
	return ret
}

func (o Id) ValidateValueForKeyPath(ioValue *[]Id, inKeyPath NSString, outError *[]NSError) bool {

	goSlice1 := make([]unsafe.Pointer,cap(*ioValue))
	for i := 0; i < len(*ioValue); i++ {
		goSlice1[i] = (*ioValue)[i].Ptr()
	}

	goSlice3 := make([]unsafe.Pointer,cap(*outError))
	for i := 0; i < len(*outError); i++ {
		goSlice3[i] = (*outError)[i].Ptr()
	}
	ret := (C.NSObject_inst_ValidateValueForKeyPath(o.Ptr(), unsafe.Pointer(&goSlice1[0]), inKeyPath.Ptr(), unsafe.Pointer(&goSlice3[0]))) != 0
	(*ioValue) = (*ioValue)[:cap(*ioValue)]
	for i := 0; i < len(*ioValue); i++ {
		if goSlice1[i] == nil {
			(*ioValue) = (*ioValue)[:i]
			break
		}
		(*ioValue)[i].ptr = goSlice1[i]
	}
	(*outError) = (*outError)[:cap(*outError)]
	for i := 0; i < len(*outError); i++ {
		if goSlice3[i] == nil {
			(*outError) = (*outError)[:i]
			break
		}
		(*outError)[i].ptr = goSlice3[i]
	}
	return ret
}

func (o Id) InverseForRelationshipKey(relationshipKey NSString) NSString {
	ret := NSString{}
	ret.ptr = C.NSObject_inst_InverseForRelationshipKey(o.Ptr(), relationshipKey.Ptr())
	return ret
}

func (o Id) CopyScriptingValue(value NSObject, key NSString, properties NSDictionary) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_CopyScriptingValue(o.Ptr(), value.Ptr(), key.Ptr(), properties.Ptr())
	return ret
}

func (o Id) ValueForKey(key NSString) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_ValueForKey(o.Ptr(), key.Ptr())
	return ret
}

func (o Id) AddObserver(observer NSObject, keyPath NSString, options NSKeyValueObservingOptions, context unsafe.Pointer)  {
	C.NSObject_inst_AddObserver(o.Ptr(), observer.Ptr(), keyPath.Ptr(), (C.NSKeyValueObservingOptions)(options), unsafe.Pointer(context))
}

func (o Id) ClassDescription() NSClassDescription {
	ret := NSClassDescription{}
	ret.ptr = C.NSObject_inst_ClassDescription(o.Ptr())
	return ret
}

func (o Id) ValueAtIndex(index NSUInteger, key NSString) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_ValueAtIndex(o.Ptr(), (C.NSUInteger)(index), key.Ptr())
	return ret
}

func (o Id) MutableOrderedSetValueForKeyPath(keyPath NSString) NSMutableOrderedSet {
	ret := NSMutableOrderedSet{}
	ret.ptr = C.NSObject_inst_MutableOrderedSetValueForKeyPath(o.Ptr(), keyPath.Ptr())
	return ret
}

func (o Id) DidChange(changeKind NSKeyValueChange, indexes NSIndexSet, key NSString)  {
	C.NSObject_inst_DidChange(o.Ptr(), (C.NSKeyValueChange)(changeKind), indexes.Ptr(), key.Ptr())
}

func (o Id) IsGreaterThanOrEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_IsGreaterThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) MutableOrderedSetValueForKey(key NSString) NSMutableOrderedSet {
	ret := NSMutableOrderedSet{}
	ret.ptr = C.NSObject_inst_MutableOrderedSetValueForKey(o.Ptr(), key.Ptr())
	return ret
}

func (o Id) PerformSelectorWithObject(aSelector SEL, anArgument NSObject, delay NSTimeInterval)  {
	C.NSObject_inst_PerformSelectorWithObject(o.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr(), (C.NSTimeInterval)(delay))
}

func (o Id) PerformSelectorWithObjectAfterDelay(aSelector SEL, anArgument NSObject, delay NSTimeInterval, modes NSArray)  {
	C.NSObject_inst_PerformSelectorWithObjectAfterDelay(o.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr(), (C.NSTimeInterval)(delay), modes.Ptr())
}

func (o Id) PerformSelectorOnThread(aSelector SEL, thr NSThread, arg NSObject, wait BOOL)  {
	C.NSObject_inst_PerformSelectorOnThread(o.Ptr(), unsafe.Pointer(aSelector), thr.Ptr(), arg.Ptr(), (C.BOOL)(wait))
}

func (o Id) PerformSelectorOnThreadWithObject(aSelector SEL, thr NSThread, arg NSObject, wait BOOL, array NSArray)  {
	C.NSObject_inst_PerformSelectorOnThreadWithObject(o.Ptr(), unsafe.Pointer(aSelector), thr.Ptr(), arg.Ptr(), (C.BOOL)(wait), array.Ptr())
}

func (o Id) IsCaseInsensitiveLike(object NSString) bool {
	ret := (C.NSObject_inst_IsCaseInsensitiveLike(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) ClassCode() FourCharCode {
	ret := (FourCharCode)(C.NSObject_inst_ClassCode(o.Ptr()))
	return ret
}

func (o Id) MutableSetValueForKey(key NSString) NSMutableSet {
	ret := NSMutableSet{}
	ret.ptr = C.NSObject_inst_MutableSetValueForKey(o.Ptr(), key.Ptr())
	return ret
}

func (o Id) SetValuesForKeysWithDictionary(keyedValues NSDictionary)  {
	C.NSObject_inst_SetValuesForKeysWithDictionary(o.Ptr(), keyedValues.Ptr())
}

func (o Id) RemoveObserverForKeyPath(observer NSObject, keyPath NSString)  {
	C.NSObject_inst_RemoveObserverForKeyPath(o.Ptr(), observer.Ptr(), keyPath.Ptr())
}

func (o Id) RemoveObserverForKeyPathContext(observer NSObject, keyPath NSString, context unsafe.Pointer)  {
	C.NSObject_inst_RemoveObserverForKeyPathContext(o.Ptr(), observer.Ptr(), keyPath.Ptr(), unsafe.Pointer(context))
}

func (o Id) ClassName() NSString {
	ret := NSString{}
	ret.ptr = C.NSObject_inst_ClassName(o.Ptr())
	return ret
}

func (o Id) ForwardInvocation(anInvocation NSInvocation)  {
	C.NSObject_inst_ForwardInvocation(o.Ptr(), anInvocation.Ptr())
}

func (o Id) ClassForKeyedArchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_ClassForKeyedArchiver(o.Ptr())))
	return ret
}

func (o Id) ToManyRelationshipKeys() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSObject_inst_ToManyRelationshipKeys(o.Ptr())
	return ret
}

func (o Id) ObjectSpecifier() NSScriptObjectSpecifier {
	ret := NSScriptObjectSpecifier{}
	ret.ptr = C.NSObject_inst_ObjectSpecifier(o.Ptr())
	return ret
}

func (o Id) MethodSignatureForSelector(aSelector SEL) NSMethodSignature {
	ret := NSMethodSignature{}
	ret.ptr = C.NSObject_inst_MethodSignatureForSelector(o.Ptr(), unsafe.Pointer(aSelector))
	return ret
}

func (o Id) ScriptingIsLessThanOrEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingIsLessThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) Copy() Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_Copy(o.Ptr())
	return ret
}

func (o Id) ReplacementObjectForCoder(aCoder NSCoder) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_ReplacementObjectForCoder(o.Ptr(), aCoder.Ptr())
	return ret
}

func (o Id) PerformSelectorInBackground(aSelector SEL, arg NSObject)  {
	C.NSObject_inst_PerformSelectorInBackground(o.Ptr(), unsafe.Pointer(aSelector), arg.Ptr())
}

func (o Id) ScriptingIsGreaterThanOrEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingIsGreaterThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) ScriptingIsGreaterThan(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingIsGreaterThan(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) Dealloc()  {
	C.NSObject_inst_Dealloc(o.Ptr())
}

func (o Id) ScriptingEndsWith(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingEndsWith(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) ToOneRelationshipKeys() NSArray {
	ret := NSArray{}
	ret.ptr = C.NSObject_inst_ToOneRelationshipKeys(o.Ptr())
	return ret
}

func (o Id) IsLike(object NSString) bool {
	ret := (C.NSObject_inst_IsLike(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) PerformSelectorOnMainThreadWithObject(aSelector SEL, arg NSObject, wait BOOL)  {
	C.NSObject_inst_PerformSelectorOnMainThreadWithObject(o.Ptr(), unsafe.Pointer(aSelector), arg.Ptr(), (C.BOOL)(wait))
}

func (o Id) PerformSelectorOnMainThreadWithObjectWaitUntilDone(aSelector SEL, arg NSObject, wait BOOL, array NSArray)  {
	C.NSObject_inst_PerformSelectorOnMainThreadWithObjectWaitUntilDone(o.Ptr(), unsafe.Pointer(aSelector), arg.Ptr(), (C.BOOL)(wait), array.Ptr())
}

func (o Id) ForwardingTargetForSelector(aSelector SEL) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_ForwardingTargetForSelector(o.Ptr(), unsafe.Pointer(aSelector))
	return ret
}

func (o Id) SetNilValueForKey(key NSString)  {
	C.NSObject_inst_SetNilValueForKey(o.Ptr(), key.Ptr())
}

func (o Id) IndicesOfObjectsByEvaluatingObjectSpecifier(specifier NSScriptObjectSpecifier) NSArray {
	ret := NSArray{}
	ret.ptr = C.NSObject_inst_IndicesOfObjectsByEvaluatingObjectSpecifier(o.Ptr(), specifier.Ptr())
	return ret
}

func (o Id) ReplacementObjectForKeyedArchiver(archiver NSKeyedArchiver) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_ReplacementObjectForKeyedArchiver(o.Ptr(), archiver.Ptr())
	return ret
}

func (o Id) ScriptingContains(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingContains(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) IsLessThan(object NSObject) bool {
	ret := (C.NSObject_inst_IsLessThan(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) Init() Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_Init(o.Ptr())
	return ret
}

func (o Id) RemoveValueAtIndex(index NSUInteger, key NSString)  {
	C.NSObject_inst_RemoveValueAtIndex(o.Ptr(), (C.NSUInteger)(index), key.Ptr())
}

func (o Id) WillChangeValueForKey(key NSString)  {
	C.NSObject_inst_WillChangeValueForKey(o.Ptr(), key.Ptr())
}

func (o Id) WillChangeValueForKeyWithSetMutation(key NSString, mutationKind NSKeyValueSetMutationKind, objects NSSet)  {
	C.NSObject_inst_WillChangeValueForKeyWithSetMutation(o.Ptr(), key.Ptr(), (C.NSKeyValueSetMutationKind)(mutationKind), objects.Ptr())
}

func (o Id) ValueForUndefinedKey(key NSString) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_ValueForUndefinedKey(o.Ptr(), key.Ptr())
	return ret
}

func (o Id) ReplaceValueAtIndex(index NSUInteger, key NSString, value NSObject)  {
	C.NSObject_inst_ReplaceValueAtIndex(o.Ptr(), (C.NSUInteger)(index), key.Ptr(), value.Ptr())
}

func (o Id) ScriptingProperties() NSDictionary {
	ret := NSDictionary{}
	ret.ptr = C.NSObject_inst_ScriptingProperties(o.Ptr())
	return ret
}

func (o Id) ValueWithName(name NSString, key NSString) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_ValueWithName(o.Ptr(), name.Ptr(), key.Ptr())
	return ret
}

func (o Id) AwakeAfterUsingCoder(aDecoder NSCoder) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_AwakeAfterUsingCoder(o.Ptr(), aDecoder.Ptr())
	return ret
}

func (o Id) SetScriptingProperties(scriptingProperties NSDictionary)  {
	C.NSObject_inst_SetScriptingProperties(o.Ptr(), scriptingProperties.Ptr())
}

func (o Id) ClassForCoder() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_ClassForCoder(o.Ptr())))
	return ret
}

func (o Id) ValueWithUniqueID(uniqueID NSObject, key NSString) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_ValueWithUniqueID(o.Ptr(), uniqueID.Ptr(), key.Ptr())
	return ret
}

func (o Id) ScriptingIsLessThan(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingIsLessThan(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) IsEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_IsEqualTo(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) SetObservationInfo(observationInfo unsafe.Pointer)  {
	C.NSObject_inst_SetObservationInfo(o.Ptr(), unsafe.Pointer(observationInfo))
}

func (o Id) IsGreaterThan(object NSObject) bool {
	ret := (C.NSObject_inst_IsGreaterThan(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) MutableArrayValueForKey(key NSString) NSMutableArray {
	ret := NSMutableArray{}
	ret.ptr = C.NSObject_inst_MutableArrayValueForKey(o.Ptr(), key.Ptr())
	return ret
}

func (o Id) DoesContain(object NSObject) bool {
	ret := (C.NSObject_inst_DoesContain(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) NewScriptingObjectOfClass(objectClass Class, key NSString, contentsValue NSObject, properties NSDictionary) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_NewScriptingObjectOfClass(o.Ptr(), unsafe.Pointer(objectClass), key.Ptr(), contentsValue.Ptr(), properties.Ptr())
	return ret
}

func (o Id) ScriptingIsEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingIsEqualTo(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) IsNotEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_IsNotEqualTo(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) ScriptingValueForSpecifier(objectSpecifier NSScriptObjectSpecifier) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_ScriptingValueForSpecifier(o.Ptr(), objectSpecifier.Ptr())
	return ret
}

func (o Id) ClassForArchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_ClassForArchiver(o.Ptr())))
	return ret
}

func (o Id) ConformsToProtocol(aProtocol Protocol) bool {
	ret := (C.NSObject_inst_ConformsToProtocol(o.Ptr(), aProtocol.Ptr())) != 0
	return ret
}

func (o Id) Description() NSString {
	ret := NSString{}
	ret.ptr = C.NSObject_inst_Description(o.Ptr())
	return ret
}

func (o Id) DebugDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSObject_inst_DebugDescription(o.Ptr())
	return ret
}

func (o Id) Hash() NSUInteger {
	ret := (NSUInteger)(C.NSObject_inst_Hash(o.Ptr()))
	return ret
}

func (o Id) GetClass() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_Class(o.Ptr())))
	return ret
}

func (o Id) Self() Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_Self(o.Ptr())
	return ret
}

func (o Id) IsProxy() bool {
	ret := (C.NSObject_inst_IsProxy(o.Ptr())) != 0
	return ret
}

func (o Id) IsMemberOfClass(aClass Class) bool {
	ret := (C.NSObject_inst_IsMemberOfClass(o.Ptr(), unsafe.Pointer(aClass))) != 0
	return ret
}

func (o Id) RespondsToSelector(aSelector SEL) bool {
	ret := (C.NSObject_inst_RespondsToSelector(o.Ptr(), unsafe.Pointer(aSelector))) != 0
	return ret
}

func (o Id) IsEqual(object NSObject) bool {
	ret := (C.NSObject_inst_IsEqual(o.Ptr(), object.Ptr())) != 0
	return ret
}

func (o Id) PerformSelector(aSelector SEL) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_PerformSelector(o.Ptr(), unsafe.Pointer(aSelector))
	return ret
}

func (o Id) PerformSelectorWithObjectWithObject(aSelector SEL, object1 NSObject, object2 NSObject) Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_PerformSelectorWithObjectWithObject(o.Ptr(), unsafe.Pointer(aSelector), object1.Ptr(), object2.Ptr())
	return ret
}

func (o Id) Retain() Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_Retain(o.Ptr())
	return ret
}

func (o Id) Release()  {
	C.NSObject_inst_Release(o.Ptr())
}

func (o Id) IsKindOfClass(aClass Class) bool {
	ret := (C.NSObject_inst_IsKindOfClass(o.Ptr(), unsafe.Pointer(aClass))) != 0
	return ret
}

func (o Id) Autorelease() Id {
	ret := Id{}
	ret.ptr = C.NSObject_inst_Autorelease(o.Ptr())
	return ret
}

func (o Id) RetainCount() NSUInteger {
	ret := (NSUInteger)(C.NSObject_inst_RetainCount(o.Ptr()))
	return ret
}

func (o Id) Zone() *_NSZone {
	ret := (*_NSZone)(unsafe.Pointer(C.NSObject_inst_Zone(o.Ptr())))
	return ret
}

func (o Id) Superclass() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_Superclass(o.Ptr())))
	return ret
}

func NSAutoreleasePoolConformsToProtocol(protocol Protocol) bool {
	ret := (C.NSAutoreleasePool_ConformsToProtocol(protocol.Ptr())) != 0
	return ret
}

func NSAutoreleasePoolVersion() NSInteger {
	ret := (NSInteger)(C.NSAutoreleasePool_Version())
	return ret
}

func NSAutoreleasePoolNew() NSAutoreleasePool {
	ret := NSAutoreleasePool{}
	ret.ptr = C.NSAutoreleasePool_New()
	return ret
}

func NSAutoreleasePoolResolveClassMethod(sel SEL) bool {
	ret := (C.NSAutoreleasePool_ResolveClassMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSAutoreleasePoolAllocWithZone(zone *_NSZone) NSAutoreleasePool {
	ret := NSAutoreleasePool{}
	ret.ptr = C.NSAutoreleasePool_AllocWithZone(unsafe.Pointer(zone))
	return ret
}

func NSAutoreleasePoolAddObject(anObject NSObject)  {
	C.NSAutoreleasePool_AddObject(anObject.Ptr())
}

func NSAutoreleasePoolDebugDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSAutoreleasePool_DebugDescription()
	return ret
}

func NSAutoreleasePoolDescription() NSString {
	ret := NSString{}
	ret.ptr = C.NSAutoreleasePool_Description()
	return ret
}

func NSAutoreleasePoolSetVersion(aVersion NSInteger)  {
	C.NSAutoreleasePool_SetVersion((C.NSInteger)(aVersion))
}

func NSAutoreleasePoolCopyWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSAutoreleasePool_CopyWithZone(unsafe.Pointer(zone))
	return ret
}

func NSAutoreleasePoolMutableCopyWithZone(zone *_NSZone) Id {
	ret := Id{}
	ret.ptr = C.NSAutoreleasePool_MutableCopyWithZone(unsafe.Pointer(zone))
	return ret
}

func NSAutoreleasePoolIsSubclassOfClass(aClass Class) bool {
	ret := (C.NSAutoreleasePool_IsSubclassOfClass(unsafe.Pointer(aClass))) != 0
	return ret
}

func NSAutoreleasePoolInstanceMethodSignatureForSelector(aSelector SEL) NSMethodSignature {
	ret := NSMethodSignature{}
	ret.ptr = C.NSAutoreleasePool_InstanceMethodSignatureForSelector(unsafe.Pointer(aSelector))
	return ret
}

func NSAutoreleasePoolHash() NSUInteger {
	ret := (NSUInteger)(C.NSAutoreleasePool_Hash())
	return ret
}

func NSAutoreleasePoolSuperclass() Class {
	ret := (Class)(unsafe.Pointer(C.NSAutoreleasePool_Superclass()))
	return ret
}

func NSAutoreleasePoolResolveInstanceMethod(sel SEL) bool {
	ret := (C.NSAutoreleasePool_ResolveInstanceMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSAutoreleasePoolInstancesRespondToSelector(aSelector SEL) bool {
	ret := (C.NSAutoreleasePool_InstancesRespondToSelector(unsafe.Pointer(aSelector))) != 0
	return ret
}

func NSAutoreleasePoolClass() Class {
	ret := (Class)(unsafe.Pointer(C.NSAutoreleasePool_Class()))
	return ret
}

func NSAutoreleasePoolLoad()  {
	C.NSAutoreleasePool_Load()
}

func NSAutoreleasePoolAlloc() NSAutoreleasePool {
	ret := NSAutoreleasePool{}
	ret.ptr = C.NSAutoreleasePool_Alloc()
	return ret
}

func (o NSAutoreleasePool) AddObject(anObject NSObject)  {
	C.NSAutoreleasePool_inst_AddObject(o.Ptr(), anObject.Ptr())
}

func (o NSAutoreleasePool) Drain()  {
	C.NSAutoreleasePool_inst_Drain(o.Ptr())
}

func MyClassAlloc() MyClass {
	ret := MyClass{}
	ret.ptr = unsafe.Pointer(C.MyClassAlloc())
	return ret
}

type MyClassDispatch struct {
	Dealloc func()
	Release func(MyClassSupermethods)
}
var MyClassLookup map[unsafe.Pointer]MyClassDispatch =
	map[unsafe.Pointer]MyClassDispatch{}

type MyClassSupermethods struct {
	Release func()

}

func (d MyClass) DeallocCallback(f func()) {
	dispatch := MyClassLookup[d.Ptr()]
	dispatch.Dealloc = f
	MyClassLookup[d.Ptr()] = dispatch
}

func (d MyClass) ReleaseCallback(f func(MyClassSupermethods)) {
	dispatch := MyClassLookup[d.Ptr()]
	dispatch.Release = f
	MyClassLookup[d.Ptr()] = dispatch
}

func (o MyClass) SuperRelease()  {
	C.MyClass_super_release(o.Ptr())
}
